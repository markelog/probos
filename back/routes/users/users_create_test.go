package users_test

import (
	"net/http"
	"testing"

	"github.com/markelog/pilgrima/test/request"
	"github.com/markelog/pilgrima/test/schema"
)

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
