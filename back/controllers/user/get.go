package users

import "github.com/markelog/probos/back/database/models"

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
