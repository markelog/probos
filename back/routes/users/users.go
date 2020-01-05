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

// Up project route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	ctrl := controller.New(db)

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
