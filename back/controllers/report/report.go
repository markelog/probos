package reports

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
)

// Report type
type Report struct {
	db    *gorm.DB
	model *gorm.DB
}

// CreateArgs are create arguments for report type
type CreateArgs struct {
	Project struct {
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
	} `json:"project"`
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
		project models.Project
		branch  models.Branch
		commit  = &models.Commit{
			BranchID: branch.ID,
			Hash:     args.Project.Branch.Commit.Hash,
			Author:   args.Project.Branch.Commit.Author,
			Message:  args.Project.Branch.Commit.Message,
			Date:     args.Project.Branch.Commit.Date,
		}

		tx = report.db.Begin()
	)

	err = tx.Where(models.Project{
		Repository: args.Project.Repository,
	}).FirstOrCreate(&project).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Branch{
		ProjectID: project.ID,
		Name:      args.Project.Branch.Name,
	}).FirstOrCreate(&branch).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Commit{
		BranchID: branch.ID,
		Hash:     args.Project.Branch.Commit.Hash,
	}).Assign(models.Commit{
		Author:  args.Project.Branch.Commit.Author,
		Message: args.Project.Branch.Commit.Message,
		Date:    args.Project.Branch.Commit.Date,
	}).FirstOrCreate(&commit).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	reports := []*models.Report{}
	for key, data := range args.Project.Branch.Commit.Report {
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

// GetArgs are arguments for get handler
type GetArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

// GetCommit is commit result to return
type GetCommit struct {
	Hash    string `json:"hash"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

// GetData is series result to return
type GetData struct {
	X      string    `json:"x"`
	Y      uint      `json:"y"`
	Commit GetCommit `json:"commit"`
}

// GetSize is size result to return
type GetSize struct {
	ID   string    `json:"id"`
	Data []GetData `json:"data"`
}

// GetResult is a return value for Get handler
type GetResult struct {
	Name  string    `json:"name"`
	Sizes []GetSize `json:"sizes"`
}

// Get reports
func (report *Report) Get(args *GetArgs) (result []*GetResult, err error) {
	var (
		commits []models.Commit

		project = report.db.Table("projects").Select("id").Where(
			"repository = ?",
			args.Repository,
		).QueryExpr()

		branch = report.db.Table("branches").Select("id").Where(
			"name = ? AND project_id = (?)",
			args.Branch, project,
		).QueryExpr()
	)

	err = report.db.Preload("Reports", func(db *gorm.DB) *gorm.DB {
		return report.db.Select("name, size, gzip, commit_id, updated_at")
	}).Where("branch_id = (?)", branch).
		Order("created_at DESC").
		Find(&commits).
		Error

	if err != nil {
		return nil, err
	}

	tmp := map[string]*GetResult{}

	// Format
	for _, commit := range commits {
		for _, report := range commit.Reports {
			if _, ok := tmp[report.Name]; !ok {
				tmp[report.Name] = &GetResult{
					Name: report.Name,
					Sizes: []GetSize{
						GetSize{
							ID:   "zip",
							Data: []GetData{},
						},
						GetSize{
							ID:   "gzip",
							Data: []GetData{},
						},
					},
				}
			}

			size := &(tmp[report.Name].Sizes[0])
			gzip := &(tmp[report.Name].Sizes[1])

			getCommit := &GetCommit{
				Hash:    commit.Hash,
				Author:  commit.Author,
				Message: commit.Message,
			}

			x := commit.Date.Format("2006-01-02T15:04")

			size.Data = append(size.Data, GetData{
				X:      x,
				Y:      report.Size,
				Commit: *getCommit,
			})
			gzip.Data = append(gzip.Data, GetData{
				X:      x,
				Y:      report.Gzip,
				Commit: *getCommit,
			})
		}
	}

	for _, data := range tmp {
		result = append(result, data)
	}

	return result, nil
}
