package users

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// RepositoriesResult is the result argument for the Repositories handler
type RepositoriesResult struct {
	Name       string `json:"name,omitempty"`
	Repository string `json:"repository,omitempty"`
}

// Repositories returns list of repos belongs to user
func (user *User) Repositories(
	username string, page int,
) ([]RepositoriesResult, error) {
	var getUser models.User
	var repos = []RepositoriesResult{}
	var correctedPage = page - 1
	var limit = 10

	reposExpr := user.db.Table("branches").
		Select("repository_id").
		QueryExpr()

	err := user.db.Where(models.User{
		Username: username,
	}).Preload("Repositories", func(db *gorm.DB) *gorm.DB {
		return user.db.Where("id IN (?)", reposExpr).
			Limit(limit).
			Offset(limit * correctedPage)
	}).Take(&getUser).Error
	if err != nil {
		return nil, err
	}

	for _, repo := range getUser.Repositories {
		result := RepositoriesResult{
			Name:       repo.Name,
			Repository: repo.Repository,
		}

		repos = append(repos, result)
	}

	return repos, nil
}
