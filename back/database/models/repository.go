package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Repository model
type Repository struct {
	gorm.Model
	Name          string `gorm:"not null;" json:"name,omitempty"`
	Repository    string `gorm:"unique; not null;" json:"repository,omitempty"`
	DefaultBranch string `json:"branch,omitempty"`
	Token         *Token
	Branches      []Branch `json:"branches,omitempty"`
	Users         []User   `gorm:"many2many:user_repositories;"`
}

var repositorySchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"repository": {"type": "string", "minLength": 1},
		"name": {"type": "string", "minLength": 1},
		"branches": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		},
		"users": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		}
	},
	"required": ["repository"]
}`)

// Validate model
func (repository Repository) Validate(db *gorm.DB) {
	repositoryLoader := gojsonschema.NewGoLoader(repository)

	check, err := gojsonschema.Validate(repositorySchema, repositoryLoader)
	if err != nil {
		db.AddError(err)
		return
	}

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
