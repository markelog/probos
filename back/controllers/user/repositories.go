package users

import (
	"time"

	"github.com/jinzhu/gorm"
	reports "github.com/markelog/probos/back/controllers/report"
	"github.com/markelog/probos/back/database/models"
)

// LastReport is a last report representation in the repo
type LastReport struct {
	Name string `json:"name"`
	Size uint   `json:"size"`
	Gzip uint   `json:"gzip"`
}

// TotalResult is the result for total size and gzip
type TotalResult struct {
	Size uint       `json:"size,omitempty"`
	Gzip uint       `json:"gzip,omitempty"`
	Date *time.Time `json:"date,omitempty"`
}

// RepositoriesResult is the result argument for the Repositories handler
type RepositoriesResult struct {
	Name       string               `json:"name,omitempty"`
	Repository string               `json:"repository,omitempty"`
	LastReport []*reports.GetResult `json:"last-report,omitempty"`
	Total      []*TotalResult       `json:"total,omitempty"`
}

// Repositories returns list of repos belongs to user
// TODO: refactor
func (user *User) Repositories(
	username string, page int,
) ([]*RepositoriesResult, error) {
	var getUser models.User
	var getTotal = []*TotalResult{}
	var repos = []*RepositoriesResult{}
	var correctedPage = page - 1
	var limit = 10

	reposExpr := user.db.Table("branches").
		Select("repository_id").
		QueryExpr()

	preloadRepos := func(db *gorm.DB) *gorm.DB {
		return user.db.Where("id IN (?)", reposExpr).
			Limit(limit).
			Offset(limit * correctedPage)
	}

	preloadCommits := func(db *gorm.DB) *gorm.DB {
		return user.db.
			Limit(5).
			Order("date DESC")
	}

	err := user.db.Where(models.User{
		Username: username,
	}).Preload("Repositories", preloadRepos).
		Preload("Repositories.Branches").
		Preload("Repositories.Branches.Commits", preloadCommits).
		Preload("Repositories.Branches.Commits.Reports").
		Take(&getUser).Error
	if err != nil {
		return nil, err
	}

	for _, repo := range getUser.Repositories {
		commits := []uint{}
		result := &RepositoriesResult{
			Name:       repo.Name,
			Repository: repo.Repository,
		}

		repos = append(repos, result)
		branchCommits := repo.Branches[0].Commits
		if len(branchCommits) == 0 {
			continue
		}

		for _, branchCommit := range branchCommits {
			commits = append(commits, branchCommit.ID)
		}

		result.LastReport = reports.FormatGetResult([]models.Commit{
			repo.Branches[0].Commits[0],
		})

		err = user.db.Table("reports").
			Select("SUM(reports.size) as size, SUM(reports.gzip) as gzip, date(commits.date)").
			Joins("INNER JOIN commits on commits.id in (reports.commit_id)").
			Group("commits.date").
			Where("commit_id in (?)", commits).
			Limit(5).
			Scan(&getTotal).Error
		if err != nil {
			return nil, err
		}

		result.Total = getTotal
	}

	return repos, nil
}
