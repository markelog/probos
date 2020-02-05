package reports

import (
	"context"
	"os"

	"github.com/google/go-github/v29/github"
	"github.com/markelog/probos/back/database/models"
	"golang.org/x/oauth2"
)

func fetchGithubUser(email string) (*models.Author, error) {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, email)
	if err != nil {
		return nil, err
	}

	result := &models.Author{
		Username: *user.Login,
		Avatar:   *user.AvatarURL,
		URL:      *user.HTMLURL,
		Email:    *user.Email,
	}

	return result, nil
}
