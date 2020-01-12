package users_test

import (
	"net/http"
	"testing"

	"github.com/markelog/probos/back/test/request"
	"github.com/markelog/probos/back/test/schema"
	"gopkg.in/gavv/httpexpect.v1"
)

func prepareTestForGet(t *testing.T) *httpexpect.Expect {
	req := request.Up(app, t)

	data := map[string]interface{}{
		"name":     "Oleg",
		"username": "markelog",
		"email":    "markelog@gmail.com",
		"avatar":   "test.png",
		"provider": "github",
	}

	req.POST("/users").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	return req
}

func TestGetNotFound(t *testing.T) {
	teardown()
	defer teardown()

	req := request.Up(app, t)

	user := req.GET("/users/markelog").
		Expect().
		Status(http.StatusBadRequest)

	json := user.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("message").Equal("record not found")
}

func TestGet(t *testing.T) {
	teardown()
	req := prepareTestForGet(t)
	defer teardown()

	user := req.GET("/users/markelog").
		Expect().
		Status(http.StatusOK)

	json := user.JSON()

	json.Schema(schema.Response)

	object := json.Object().Value("payload").Object()

	object.NotContainsKey("ID")
	object.NotContainsKey("CreatedAt")
	object.NotContainsKey("UpdatedAt")
	object.NotContainsKey("DeletedAt")

	object.Value("name").Equal("Oleg")
	object.Value("username").Equal("markelog")
	object.Value("email").Equal("markelog@gmail.com")
	object.Value("avatar").Equal("test.png")
	object.Value("provider").Equal("github")
}
