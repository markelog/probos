package users

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// User type
type User struct {
	db *gorm.DB
}

// CreateArgs are create arguments for user type
type CreateArgs struct {
	Name         string   `json:"name,omitempty"`
	Username     string   `json:"username,omitempty"`
	Email        string   `json:"email,omitempty"`
	Avatar       string   `json:"avatar,omitempty"`
	Provider     string   `json:"provider,omitempty"`
	Repositories []string `json:"repositories,omitempty"`
}

// New user
func New(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// Create user
func (user *User) Create(args *CreateArgs) error {
	repositories := []models.Repository{}
	for _, repository := range args.Repositories {
		prj := models.Repository{
			Repository: fmt.Sprintf("github.com/%s/", repository),
		}
		repositories = append(repositories, prj)
	}

	spew.Dump(repositories)

	data := &models.User{
		Name:         args.Name,
		Username:     args.Username,
		Email:        args.Email,
		Avatar:       args.Avatar,
		Provider:     args.Provider,
		Repositories: repositories,
	}

	err := user.db.Where(models.User{
		Username: args.Username,
	}).Assign(&data).FirstOrCreate(&data).Error

	if err != nil {
		return err
	}

	return nil
}

// GetResult is the result argument for the Get handler
type GetResult struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// Get user
func (user *User) Get(username string) (*GetResult, error) {
	var getUser models.User

	err := user.db.Where(models.User{
		Username: username,
	}).Take(&getUser).Error
	if err != nil {
		return nil, err
	}

	result := &GetResult{
		Name:     getUser.Name,
		Username: getUser.Username,
		Email:    getUser.Email,
		Avatar:   getUser.Avatar,
		Provider: getUser.Provider,
	}

	return result, nil
}
