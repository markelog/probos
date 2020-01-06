package users_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/users"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/routes"
)

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	users.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}
