package reports

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// Report type
type Report struct {
	db *gorm.DB
}

// CreateArgs are create arguments for report type
type CreateArgs struct {
	Repository struct {
		Repository    string `json:"repository"`
		DefaultBranch string `json:"defaultBranch"`
		Branch        struct {
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
	} `json:"repository"`
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
		authorEmail = args.Repository.Branch.Commit.Author
		author      *models.Author
		branch      models.Branch
		repository  models.Repository
		commit      = &models.Commit{
			BranchID: branch.ID,
			Hash:     args.Repository.Branch.Commit.Hash,
			Message:  args.Repository.Branch.Commit.Message,
			Date:     args.Repository.Branch.Commit.Date,
		}

		tx = report.db.Begin()
	)

	author, err = fetchGithubUser(authorEmail)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Repository{
		Repository: args.Repository.Repository,
	}).Assign(&models.Repository{
		DefaultBranch: args.Repository.DefaultBranch,
	}).Take(&repository).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Branch{
		RepositoryID: repository.ID,
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
		Author:  *author,
		Message: args.Repository.Branch.Commit.Message,
		Date:    args.Repository.Branch.Commit.Date,
	}).FirstOrCreate(&commit).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Author{
		Email: authorEmail,
	}).FirstOrCreate(&author).Error
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
		return errors.New("There is no applicable reports")
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
