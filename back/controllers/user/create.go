package users

import (
	"fmt"

	"github.com/markelog/probos/back/database/models"
)

// RepositoryArg is the repository argument for CreateArgs
type RepositoryArg struct {
	Name   string `json:"name,omitempty"`
	Branch string `json:"branch,omitempty"`
}

// CreateArgs are create arguments for user type
type CreateArgs struct {
	Name         string           `json:"name,omitempty"`
	Username     string           `json:"username,omitempty"`
	Email        string           `json:"email,omitempty"`
	Avatar       string           `json:"avatar,omitempty"`
	Provider     string           `json:"provider,omitempty"`
	Repositories []*RepositoryArg `json:"repositories,omitempty"`
}

// Create user
func (user *User) Create(args *CreateArgs) error {
	repositories := []models.Repository{}
	repositoriesNames := []string{}
	repositoriesBranches := []string{}
	for _, repository := range args.Repositories {
		name := fmt.Sprintf("github.com/%s", repository.Name)
		prj := models.Repository{
			Repository: name,
		}

		repositoriesNames = append(repositoriesNames, name)
		repositoriesBranches = append(repositoriesBranches, repository.Branch)
		repositories = append(repositories, prj)
	}

	data := &models.User{
		Name:     args.Name,
		Username: args.Username,
		Email:    args.Email,
		Avatar:   args.Avatar,
		Provider: args.Provider,
	}

	tx := user.db.Begin()

	err := tx.Where(models.User{
		Username: args.Username,
	}).FirstOrCreate(&data).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Table("repositories").
		Where("repository IN (?)", repositoriesNames).
		Update("default_branch", repositoriesBranches).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Find(&repositories).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&data).
		Association("Repositories").
		Replace(&repositories).Error
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
