package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/markelog/probos/back/database"
	"github.com/markelog/probos/back/env"
	"github.com/markelog/probos/back/logger"
	"github.com/markelog/probos/back/routes"
	"github.com/markelog/probos/back/routes/common"
	"github.com/markelog/probos/back/routes/github"
	"github.com/markelog/probos/back/routes/reports"
	"github.com/markelog/probos/back/routes/repositories"
	"github.com/markelog/probos/back/routes/root"
	"github.com/markelog/probos/back/routes/tokens"
	"github.com/markelog/probos/back/routes/users"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Up()

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	var (
		app = routes.Up()
		db  = database.Up()
		log = logger.Up()
	)

	defer db.Close()

	root.Up(app, db, log)
	tokens.Up(app, db, log)
	repositories.Up(app, db, log)
	reports.Up(app, db, log)
	users.Up(app, db, log)
	common.Up(app, db, log)
	github.Up(app, db, log)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
