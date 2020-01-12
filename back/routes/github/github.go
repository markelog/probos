package github

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	ctrl "github.com/markelog/probos/back/controllers/github"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/sirupsen/logrus"
)

// Up github hook route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	config := githubapp.Config{}

	id, err := strconv.Atoi(os.Getenv("GITHUB_APP_IDENTIFIER"))
	if err != nil {
		log.Panic(err)
	}

	config.App.IntegrationID = id
	config.App.PrivateKey = os.Getenv("GITHUB_APP_PRIVATE_KEY")
	config.App.WebhookSecret = os.Getenv("GITHUB_WEBHOOK_SECRET")

	cc, err := githubapp.NewDefaultCachingClientCreator(config)
	if err != nil {
		log.Panic(err)
	}

	dispatcher := githubapp.NewDefaultEventDispatcher(
		config, ctrl.New(cc, db, log),
	)

	handler := func(
		w http.ResponseWriter,
		r *http.Request,
		router http.HandlerFunc,
	) {
		path := r.URL.Path
		isHookRoute := strings.HasPrefix(path, "/github/hook")

		if isHookRoute && r.Method == "POST" {
			dispatcher.ServeHTTP(w, r)
			return
		}

		// otherwise continue serving routes as usual
		router.ServeHTTP(w, r)
	}

	app.WrapRouter(handler)
}
