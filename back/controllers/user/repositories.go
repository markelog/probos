package users

import (
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

// RepositoriesResult is the result argument for the Repositories handler
type RepositoriesResult struct {
	Name       string               `json:"name,omitempty"`
	Repository string               `json:"repository,omitempty"`
	LastReport []*reports.GetResult `json:"last-report,omitempty"`
}

// Repositories returns list of repos belongs to user
func (user *User) Repositories(
	username string, page int,
) ([]*RepositoriesResult, error) {
	var getUser models.User
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
			Limit(1).
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
		result := &RepositoriesResult{
			Name:       repo.Name,
			Repository: repo.Repository,
		}

		repos = append(repos, result)

		if len(repo.Branches[0].Commits) == 0 {
			continue
		}

		result.LastReport = reports.FormatGetResult([]models.Commit{
			repo.Branches[0].Commits[0],
		})
	}

	return repos, nil
}
