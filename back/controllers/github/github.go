package github

import (
	"context"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/github"
	"github.com/palantir/go-githubapp/githubapp"
)

type Installation struct {
	githubapp.ClientCreator
}

func (installation *Installation) Handles() []string {
	return []string{"integration_installation"}
}

func (installation *Installation) Handle(ctx context.Context, eventType, deliveryID string, payload []byte) error {
	spew.Dump(321)

	// from github.com/google/go-github/github
	var event github.InstallationEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	spew.Dump(event)

	return nil
}
