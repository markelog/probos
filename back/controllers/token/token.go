package token

import (
	"crypto/rand"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/database/models"
)

// Token type
type Token struct {
	Token      string
	Repository uint
	db         *gorm.DB
	Model      *models.Token
}

func generate() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)

	return fmt.Sprintf("%x", bytes)
}

// New Token
func New(db *gorm.DB) Token {
	generated := generate()

	return Token{
		db: db,
		Model: &models.Token{
			Token: generated,
		},
	}
}

// Create token
func (token *Token) Create(Repository uint) (*models.Token, error) {
	var (
		RepositoryModel models.Repository
		value           = token.db.Model(token.Repository).First(&RepositoryModel)
	)

	if value.Error != nil {
		return nil, value.Error
	}

	token.Model = &models.Token{
		Token: token.Token,
	}

	value = token.db.Create(&token.Model)
	if value.Error != nil {
		return nil, value.Error
	}

	return token.Model, nil
}
