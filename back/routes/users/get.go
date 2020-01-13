package users

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/probos/back/controllers/user"
	"github.com/sirupsen/logrus"
)

func initGetRoutes(
	app *iris.Application,
	db *gorm.DB,
	log *logrus.Logger,
	ctrl *controller.User,
) {
	app.Get("/users/{username:string}", func(ctx iris.Context) {
		username := ctx.Params().Get("username")

		user, err := ctrl.Get(username)
		if err != nil {
			setGetError(log, username, ctx, err)
			return
		}

		log.WithFields(logrus.Fields{
			"username": username,
		}).Info("User is returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": user,
		})
	})

	app.Get("/users/{username:string}/repos", func(ctx iris.Context) {
		username := ctx.Params().Get("username")
		page, _ := ctx.URLParamInt("page")

		repositories, err := ctrl.Repositories(username, page)
		if err != nil {
			setGetError(log, username, ctx, err)
			return
		}

		log.WithFields(logrus.Fields{
			"username": username,
			"branch":   "master",
			"page":     page,
		}).Info("Repositories are returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": repositories,
		})
	})

}
