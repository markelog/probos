package reports

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// GetArgs are arguments for get handler
type GetArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

// GetAuthor is the struct for author data
type GetAuthor struct {
	Username string `json:"username"`
	URL      string `json:"url"`
	Avatar   string `json:"avatar"`
}

// GetSize is size result to return
type GetSize struct {
	Hash    string    `json:"hash"`
	Author  GetAuthor `json:"author"`
	Message string    `json:"message"`
	Date    string    `json:"date"`
	Size    uint      `json:"size"`
	Gzip    uint      `json:"gzip"`
}

// SizesResult is the result for the sizes
type SizesResult struct {
	Name  string    `json:"name"`
	Sizes []GetSize `json:"sizes"`
}

// GetResult is a return value for Get handler
type GetResult struct {
	Name  string         `json:"name"`
	Sizes []*SizesResult `json:"sizes"`
}

// Get reports
func (report *Report) Get(args *GetArgs) (*GetResult, error) {
	var (
		commits    []models.Commit
		repository models.Repository

		Repository = report.db.Table("repositories").Select("id").Where(
			"repository = ?",
			args.Repository,
		).QueryExpr()

		branch = report.db.Table("branches").Select("id").Where(
			"name = ? AND repository_id = (?)",
			args.Branch, Repository,
		).QueryExpr()
	)

	err := report.db.Preload("Reports", func(db *gorm.DB) *gorm.DB {
		return report.db.Select("name, size, gzip, commit_id, updated_at")
	}).Preload("Author").
		Where("branch_id = (?)", branch).
		Order("date ASC").
		Find(&commits).
		Error

	if err != nil {
		return nil, err
	}

	err = report.db.Where(models.Repository{
		Repository: args.Repository,
	}).Take(&repository).Error
	if err != nil {
		return nil, err
	}

	result := &GetResult{
		Name:  repository.Name,
		Sizes: FormatGetResult(commits),
	}

	return result, nil
}

// FormatGetResult format the result in the nice way
func FormatGetResult(commits []models.Commit) []*SizesResult {
	var result []*SizesResult
	var tmpKeys = []string{}
	var tmp = map[string]*SizesResult{}

	for _, commit := range commits {
		for _, report := range commit.Reports {
			if _, ok := tmp[report.Name]; !ok {
				tmpKeys = append(tmpKeys, report.Name)
				tmp[report.Name] = &SizesResult{
					Name: report.Name,
				}
			}

			getSize := GetSize{
				Hash: commit.Hash,
				Author: GetAuthor{
					Username: commit.Author.Username,
					Avatar:   commit.Author.Avatar,
					URL:      commit.Author.URL,
				},
				Message: commit.Message,
				Date:    commit.Date.Format("2006-01-02T15:04"),
				Size:    report.Size,
				Gzip:    report.Gzip,
			}

			tmp[report.Name].Sizes = append(tmp[report.Name].Sizes, getSize)
		}
	}

	for _, name := range tmpKeys {
		result = append(result, tmp[name])
	}

	return result
}

// LastArgs are arguments for Last handler
type LastArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

type lastValue struct {
	Size uint `json:"size"`
	Gzip uint `json:"gzip"`
}

// LastResult is a return value for Last handler
type LastResult map[string]lastValue

// Last will get you last report
func (report *Report) Last(args *LastArgs) (result LastResult, err error) {
	var (
		reports []models.Report

		project = report.db.Table("projects").Select("id").Where(
			"repository = ?",
			args.Repository,
		).QueryExpr()

		branch = report.db.Table("branches").Select("id").Where(
			"name = ? AND project_id = (?)",
			args.Branch, project,
		).QueryExpr()

		commit = report.db.Table("commits").Select("id").Where(
			"branch_id = (?)",
			branch,
		).Order("created_at DESC").Limit(1).QueryExpr()
	)

	err = report.db.Select("DISTINCT(name), size, gzip").Where(
		"commit_id = (?)",
		commit,
	).Find(&reports).Error

	if err != nil {
		return nil, err
	}

	result = make(map[string]lastValue)
	for _, report := range reports {
		result[report.Name] = lastValue{
			Size: report.Size,
			Gzip: report.Gzip,
		}
	}

	return result, err
}
