package repositories_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/probos/back/database"
	"github.com/markelog/probos/back/logger"
	"github.com/markelog/probos/back/routes/repositories"
	"github.com/markelog/probos/back/test/env"
	"github.com/markelog/probos/back/test/request"
	"github.com/markelog/probos/back/test/routes"
	"github.com/markelog/probos/back/test/schema"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Raw("TRUNCATE users CASCADE;").Row()
	db.Raw("TRUNCATE Repositories CASCADE;").Row()
	db.Raw("TRUNCATE tokens CASCADE;").Row()
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	Repositories.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestAbsenceOfARepository(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name": "test",
	}

	token := req.POST("/Repositories").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("repository: String length must be greater than or equal to 1")
}

func TestAbsenceOfAName(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"repository": "https://github.com/markelog/probos",
	}

	token := req.POST("/Repositories").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("name: String length must be greater than or equal to 1")
}

func TestAbsence(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	token := req.POST("/Repositories").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusBadRequest)

	json := token.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Contains(
			"name: String length must be greater than or equal to 1",
		)
}

func TestSuccess(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":       "yo",
		"repository": "github.com/markelog/probos",
	}

	Repository := req.POST("/Repositories").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	Repository.JSON().Schema(schema.Response)
}

func TestList(t *testing.T) {
	defer teardown()
	teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":       "yo",
		"repository": "github.com/markelog/probos",
	}

	req.POST("/Repositories").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	result := req.GET("/Repositories").
		Expect().
		Status(http.StatusOK).
		JSON()

	result.Schema(schema.Response)

	element := result.Object().Value("payload").Array().
		Element(0).Object()

	element.Value("name").Equal("yo")
	element.Value("repository").Equal("github.com/markelog/probos")
}

func TestAbsentList(t *testing.T) {
	defer teardown()
	teardown()
	req := request.Up(app, t)

	response := req.GET("/Repositories").
		Expect().
		Status(http.StatusNotFound)

	json := response.JSON()

	json.Schema(schema.Response)

	json.Schema(schema.Response)
	json.Object().Value("payload").Object().Empty()
	json.Object().Value("message").Equal("Not found")
	json.Object().Value("status").Equal("failed")
}
