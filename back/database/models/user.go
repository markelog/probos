package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// User model
type User struct {
	gorm.Model
	Name         string       `gorm:"not null;" json:"name,omitempty"`
	Username     string       `gorm:"not null;" json:"username,omitempty"`
	Email        string       `gorm:"not null;" json:"email,omitempty"`
	Avatar       string       `gorm:"not null;" json:"avatar,omitempty"`
	Provider     string       `gorm:"not null;" json:"provider,omitempty"`
	Repositories []Repository `gorm:"many2many:user_repositories;association_autocreate:false;"`
}

var userSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"username": {"type": "string", "minLength": 1},
		"email": {"type": "string", "format": "email"},
		"avatar": {"type": "string", "minLength": 1},
		"provider": {"type": "string", "minLength": 1}
	},
	"required": ["name", "username", "email", "avatar", "provider"]
}`)

// Validate model
func (user User) Validate(db *gorm.DB) {
	userLoader := gojsonschema.NewGoLoader(user)

	check, err := gojsonschema.Validate(userSchema, userLoader)
	if err != nil {
		db.AddError(err)
		return
	}

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}

}
