package users

import (
	"github.com/jinzhu/gorm"
)

// User type
type User struct {
	db *gorm.DB
}

// New user
func New(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}
