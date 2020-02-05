package models

import (
	"time"

	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Commit model struct
type Commit struct {
	gorm.Model
	Hash     string    `gorm:"unique;not null;" json:"hash,omitempty"`
	Message  string    `json:"message,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	BranchID uint      `json:"branch,omitempty"`
	Author   Author    `json:"author,omitempty"`
	AuthorID uint
	Reports  []Report `json:"reports,omitempty"`
}

var commitSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"hash": {"type": "string", "minLength": 1},
		"message": {"type": "string", "minLength": 1},
		"date": {"type": "string", "minLength": 1},
		"branch": {"type": "number", "minimum": 1},
		"reports": {
			"type": "array", 
			"items": {
				"type": "object"
			}
		},
		"author": {
			"type": "object"
		}
	},
	"required": ["hash", "author", "date", "message", "branch"]
}`)

// Validate model
func (commit Commit) Validate(db *gorm.DB) {
	commitLoader := gojsonschema.NewGoLoader(commit)

	check, err := gojsonschema.Validate(commitSchema, commitLoader)
	if err != nil {
		db.AddError(err)
		return
	}

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
