package users_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/markelog/probos/back/database"
	"github.com/markelog/probos/back/logger"
	"github.com/markelog/probos/back/routes/users"
	"github.com/markelog/probos/back/test/env"
	"github.com/markelog/probos/back/test/routes"
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
