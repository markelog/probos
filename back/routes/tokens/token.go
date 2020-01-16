package tokens

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/probos/back/controllers/token"
	"github.com/sirupsen/logrus"
)

type postRepository struct {
	Repository uint `json:"Repository"`
}

// Up token route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	ctrl := controller.New(db)

	app.Post("/tokens", func(ctx iris.Context) {
		var params postRepository
		ctx.ReadJSON(&params)

		result, err := ctrl.Create(params.Repository)

		if err != nil && err.Error() == "record not found" {
			log.WithFields(logrus.Fields{
				"Repository": params.Repository,
			}).Error("Can't find this Repository")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't find this Repository",
				"payload": iris.Map{},
			})
			return
		}

		if err != nil {
			log.WithFields(logrus.Fields{
				"Repository": params.Repository,
				"error":      err,
			}).Error("Couldn't create the token")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Something wen't wrong",
				"payload": iris.Map{},
			})
			return
		}

		log.WithFields(logrus.Fields{
			"Repository": params.Repository,
			"token":      result.Token,
		}).Info("Token created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"token": result.Token,
			},
		})
	})
}
