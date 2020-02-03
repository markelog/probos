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
func (repository *Repository) Create(name, repo, branch string) (
	*models.Repository, error,
) {
	model := &models.Repository{
		Name:          name,
		DefaultBranch: branch,
		Repository:    repo,
		Token:         token.New(repository.db).Model,
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
