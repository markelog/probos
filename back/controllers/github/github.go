package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/jinzhu/gorm"
	"github.com/markelog/probos/back/controllers/repository"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/sirupsen/logrus"
)

// Installation is an event for github hook
type Installation struct {
	cc  githubapp.ClientCreator
	db  *gorm.DB
	log *logrus.Logger
}

// New user
func New(
	cc githubapp.ClientCreator,
	db *gorm.DB,
	log *logrus.Logger,
) *Installation {
	return &Installation{
		db:  db,
		cc:  cc,
		log: log,
	}
}

// Handles returns the slice of the events this strcuct supports
func (installation *Installation) Handles() []string {
	return []string{"integration_installation"}
}

// HandleCreateEvent handles create event
func (installation *Installation) HandleCreateEvent(
	repos []*github.Repository,
	sender *github.User,
) error {
	tx := installation.db.Begin()
	prj := repository.New(tx)

	for _, repo := range repos {
		address := fmt.Sprintf("github.com/%s", *repo.FullName)
		name := *repo.Name

		_, err := prj.Create(name, address)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	installation.log.WithFields(logrus.Fields{
		"action": "create repositories",
		"login":  *sender.Login,
	}).Info("GitHub hook was called")

	return nil
}

// Handle is a general handle for supported events
func (installation *Installation) Handle(
	ctx context.Context,
	eventType, deliveryID string,
	payload []byte,
) error {

	// from github.com/google/go-github/github
	var event github.InstallationEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	if *event.Action == "created" {
		err := installation.HandleCreateEvent(
			event.Repositories,
			event.Sender,
		)
		if err != nil {
			installation.log.WithFields(logrus.Fields{
				"action": "create repositories",
				"error":  err,
				"login":  *event.Sender.Login,
			}).Error("GitHub hook failed")
			return err
		}
	}

	return nil
}
