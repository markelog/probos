package users

import (
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

	err := user.db.Where(models.User{
		Username: username,
	}).
		Preload("Repositories", "id IN(SELECT repository_id FROM branches)").
		Take(&getUser).Error
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
