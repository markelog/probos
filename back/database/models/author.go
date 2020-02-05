package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Author model
type Author struct {
	gorm.Model
	Username string   `gorm:"not null;" json:"username,omitempty"`
	Avatar   string   `gorm:"not null;" json:"avatar,omitempty"`
	URL      string   `gorm:"not null;" json:"url,omitempty"`
	Email    string   `gorm:"not null;" json:"email,omitempty"`
	Commits  []Commit `json:"commits,omitempty"`
}

var authorSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"username": {"type": "string", "minLength": 1},
		"email": {"type": "string", "format": "email"},
		"avatar": {"type": "string", "minLength": 1},
		"url": {"type": "string", "minLength": 1}
	},
	"required": ["username", "avatar", "url", "email"]
}`)

// Validate model
func (author Author) Validate(db *gorm.DB) {
	authorLoader := gojsonschema.NewGoLoader(author)

	check, err := gojsonschema.Validate(authorSchema, authorLoader)
	if err != nil {
		db.AddError(err)
		return
	}

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}

}
