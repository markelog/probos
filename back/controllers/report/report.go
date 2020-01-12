package reports

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// Report type
type Report struct {
	db    *gorm.DB
	model *gorm.DB
}

// CreateArgs are create arguments for report type
type CreateArgs struct {
	Repository struct {
		Repository string `json:"repository"`
		Branch     struct {
			Name   string `json:"name"`
			Commit struct {
				Hash    string    `json:"hash"`
				Author  string    `json:"author"`
				Message string    `json:"message"`
				Date    time.Time `json:"date"`
				Report  map[string]struct {
					Size uint `json:"size"`
					Gzip uint `json:"gzip"`
				} `json:"report"`
			} `json:"commit"`
		} `json:"branch"`
	} `json:"Repository"`
}

// New report
func New(db *gorm.DB) *Report {
	return &Report{
		db: db,
	}
}

// Create report and associated data
func (report *Report) Create(args *CreateArgs) (err error) {
	var (
		Repository models.Repository
		branch     models.Branch
		commit     = &models.Commit{
			BranchID: branch.ID,
			Hash:     args.Repository.Branch.Commit.Hash,
			Author:   args.Repository.Branch.Commit.Author,
			Message:  args.Repository.Branch.Commit.Message,
			Date:     args.Repository.Branch.Commit.Date,
		}

		tx = report.db.Begin()
	)

	err = tx.Where(models.Repository{
		Repository: args.Repository.Repository,
	}).FirstOrCreate(&Repository).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Branch{
		RepositoryID: Repository.ID,
		Name:         args.Repository.Branch.Name,
	}).FirstOrCreate(&branch).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Commit{
		BranchID: branch.ID,
		Hash:     args.Repository.Branch.Commit.Hash,
	}).Assign(models.Commit{
		Author:  args.Repository.Branch.Commit.Author,
		Message: args.Repository.Branch.Commit.Message,
		Date:    args.Repository.Branch.Commit.Date,
	}).FirstOrCreate(&commit).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	reports := []*models.Report{}
	for key, data := range args.Repository.Branch.Commit.Report {
		reports = append(reports, &models.Report{
			Name: key,
			Size: data.Size,
			Gzip: data.Gzip,
		})
	}

	if len(reports) == 0 {
		tx.Rollback()
		return errors.New("There is not applicable reports")
	}

	err = tx.Model(&commit).Association("Reports").Append(reports).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
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

		Repository = report.db.Table("repositories").Select("id").Where(
			"repository = ?",
			args.Repository,
		).QueryExpr()

		branch = report.db.Table("branches").Select("id").Where(
			"name = ? AND repository_id = (?)",
			args.Branch, Repository,
		).QueryExpr()

		commit = report.db.Table("commits").Select("id").Where(
			"branch_id = (?)",
			branch,
		).Limit(1).QueryExpr()
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
