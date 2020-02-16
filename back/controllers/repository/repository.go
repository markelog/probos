package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/controllers/token"
	"github.com/markelog/probos/back/database/models"
)

// Repository type
type Repository struct {
	db    *gorm.DB
	Model *models.Repository
}

// New Repository
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create Repository
func (repository *Repository) Create(name, repo string) (
	*models.Repository, error,
) {
	model := &models.Repository{
		Name:       name,
		Repository: repo,
		Token:      token.New(repository.db).Model,
	}

	result := repository.db.Where(models.Repository{
		Repository: repo,
	}).FirstOrCreate(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

// ListValue result value for List() method
type ListValue struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

// List Repositories
func (repository *Repository) List() ([]ListValue, error) {
	var (
		repos  []models.Repository
		result []ListValue
	)

	err := repository.db.Find(&repos).Error
	if err != nil {
		return nil, err
	}

	for _, repo := range repos {
		result = append(result, ListValue{
			Name:       repo.Name,
			Repository: repo.Repository,
		})
	}

	return result, nil
}

// GetValue is the return value for the Get()
type GetValue struct {
	Name          string   `json:"name"`
	Repository    string   `json:"repository"`
	DefaultBranch string   `json:"defaultBranch"`
	Branches      []string `json:"branches"`
}

// Get returns the repository found by the repository path
func (repository *Repository) Get(path, user string) (*GetValue, error) {
	var (
		repo   models.Repository
		result *GetValue
	)

	err := repository.db.Where(models.Repository{
		Repository: path,
	}).Preload("Branches").Take(&repo).Error
	if err != nil {
		return nil, err
	}

	result = &GetValue{
		Name:          repo.Name,
		Repository:    repo.Repository,
		DefaultBranch: repo.DefaultBranch,
	}

	for _, branch := range repo.Branches {
		result.Branches = append(result.Branches, branch.Name)
	}

	return result, nil
}
