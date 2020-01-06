package users

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/pilgrima/controllers/user"
	"github.com/sirupsen/logrus"
)

func setPostError(log *logrus.Logger, params *controller.CreateArgs, ctx iris.Context, err error) {
	errorString := err.Error()

	log.WithFields(logrus.Fields{
		"email":    params.Email,
		"username": params.Username,
	}).Error(errorString)

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "failed",
		"message": errorString,
		"payload": iris.Map{},
	})
}

func setGetError(log *logrus.Logger, username string, ctx iris.Context, err error) {
	errorString := err.Error()

	log.WithFields(logrus.Fields{
		"username": username,
	}).Error(errorString)

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "failed",
		"message": errorString,
		"payload": iris.Map{},
	})
}

// Up project route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	ctrl := controller.New(db)

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
			"status":  "created",
			"message": "Yey!",
			"payload": user,
		})
	})

	app.Post("/users", func(ctx iris.Context) {
		var params controller.CreateArgs

		err := ctx.ReadJSON(&params)
		if err != nil {
			setPostError(log, &params, ctx, err)
			return
		}

		err = ctrl.Create(&params)
		if err != nil {
			setPostError(log, &params, ctx, err)
			return
		}

		log.WithFields(logrus.Fields{
			"username": params.Username,
			"email":    params.Email,
		}).Info("User created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{},
		})
	})
}
