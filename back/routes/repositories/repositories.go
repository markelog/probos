package repositories

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/probos/back/controllers/repository"
	"github.com/markelog/probos/back/routes/middleware"
	"github.com/sirupsen/logrus"
)

type postRepository struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

// Up Repository route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	middlewares := middleware.Up()
	ctrl := controller.New(db)

	app.Post("/repositories", func(ctx iris.Context) {
		var params postRepository
		ctx.ReadJSON(&params)
		result, err := ctrl.Create(params.Name, params.Repository)

		if err != nil {
			log.WithFields(logrus.Fields{
				"name":       params.Name,
				"repository": params.Repository,
			}).Error("Can't create the Repository")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Can't create the Repository",
				"payload": iris.Map{},
			})

			return
		}

		log.WithFields(logrus.Fields{
			"Repository": params.Name,
		}).Info("Repository created")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "created",
			"message": "Yey!",
			"payload": iris.Map{
				"Repository": result.Name,
				"id":         result.ID,
			},
		})
	})

	app.Get("/repositories", func(ctx iris.Context) {
		repositories, err := ctrl.List()

		if err != nil {
			errorString := err.Error()
			log.Error(errorString)

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Something went wrong",
				"payload": iris.Map{},
			})

			return
		}

		if len(repositories) == 0 {
			log.Error("Can't find any repositories")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": iris.Map{},
			})

			return
		}

		log.Info("repositories returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": repositories,
		})
	})

	app.Get(`/repositories/{repository:path}`, func(ctx iris.Context) {
		repo := ctx.Params().Get("repository")
		user := ""

		err := middlewares.JWT.CheckJWT(ctx)

		if err == nil {
			user = ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(string)
		}

		repository, err := ctrl.Get(repo, user)
		if err != nil {
			log.Error(err.Error())

			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Bad Auth",
				"payload": iris.Map{},
			})

			return
		}

		log.Info("repositories returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": repository,
		})
	})
}
