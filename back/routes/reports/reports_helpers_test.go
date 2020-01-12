package reports_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/probos/back/database"
	"github.com/markelog/probos/back/logger"
	"github.com/markelog/probos/back/routes/reports"
	"github.com/markelog/probos/back/test/env"
	"github.com/markelog/probos/back/test/routes"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Raw("TRUNCATE users CASCADE;").Row()
	db.Raw("TRUNCATE Repositories CASCADE;").Row()
	db.Raw("TRUNCATE reports CASCADE;").Row()
	db.Raw("TRUNCATE tokens CASCADE;").Row()
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	reports.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}
