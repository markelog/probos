package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Branch model
type Branch struct {
	gorm.Model
	Name         string   `json:"name,omitempty"`
	RepositoryID uint     `json:"repository,omitempty"`
	Commits      []Commit `json:"commits,omitempty"`
}

var branchSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"Repository": {"type": "number", "minimum": 1},
		"commits": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		}
	},
	"required": ["name", "Repository"]
}`)

// Validate model
func (branch Branch) Validate(db *gorm.DB) {
	branchLoader := gojsonschema.NewGoLoader(branch)

	check, err := gojsonschema.Validate(branchSchema, branchLoader)
	if err != nil {
		db.AddError(err)
		return
	}

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
