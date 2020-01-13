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

// GetSize is size result to return
type GetSize struct {
	Hash    string `json:"hash"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Size    uint   `json:"size"`
	Gzip    uint   `json:"gzip"`
}

// GetResult is a return value for Get handler
type GetResult struct {
	Name  string    `json:"name"`
	Sizes []GetSize `json:"sizes"`
}

// Get reports
func (report *Report) Get(args *GetArgs) ([]*GetResult, error) {
	var (
		commits []models.Commit

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
	}).Where("branch_id = (?)", branch).
		Order("date ASC").
		Find(&commits).
		Error

	if err != nil {
		return nil, err
	}

	return formatGetResult(commits), nil
}

func formatGetResult(commits []models.Commit) []*GetResult {
	var result []*GetResult
	var tmpKeys = []string{}
	var tmp = map[string]*GetResult{}

	for _, commit := range commits {
		for _, report := range commit.Reports {
			if _, ok := tmp[report.Name]; !ok {
				tmpKeys = append(tmpKeys, report.Name)
				tmp[report.Name] = &GetResult{
					Name: report.Name,
				}
			}

			getSize := GetSize{
				Hash:    commit.Hash,
				Author:  commit.Author,
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
