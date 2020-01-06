package users_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/pilgrima/database"
	"github.com/markelog/pilgrima/logger"
	"github.com/markelog/pilgrima/routes/users"
	"github.com/markelog/pilgrima/test/env"
	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/routes"
	"github.com/markelog/pilgrima/test/schema"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Raw("TRUNCATE users CASCADE;").Row()
}

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

func TestAbsenceOfUsername(t *testing.T) {
	teardown()
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":     "killa",
		"email":    "killa@gorilla.com",
		"avatar":   "test.png",
		"provider": "github",
	}

	user := req.POST("/users").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := user.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("message").Equal("(root): username is required")
}

func TestCreate(t *testing.T) {
	teardown()
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":     "killa",
		"username": "gorilla",
		"email":    "killa@gorilla.com",
		"avatar":   "test.png",
		"provider": "github",
	}

	user := req.POST("/users").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	json := user.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("status").String().
		Equal("created")
}
