package users

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
)

// User type
type User struct {
	db    *gorm.DB
	model *gorm.DB
}

// CreateArgs are create arguments for user type
type CreateArgs struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// New user
func New(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// Create user and associated data
func (user *User) Create(args *CreateArgs) error {
	data := &models.User{
		Name:     args.Name,
		Username: args.Username,
		Email:    args.Email,
		Avatar:   args.Avatar,
	}

	err := user.db.Where(models.User{
		Username: args.Username,
	}).Assign(&data).FirstOrCreate(&data).Error

	if err != nil {
		return err
	}

	return nil
}