package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/probos/back/controllers/repository"
	"github.com/sirupsen/logrus"
)

type postRepository struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

// Up Repository route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/repositories", func(ctx iris.Context) {
		var params postRepository
		ctx.ReadJSON(&params)
		ctrl := controller.New(db)
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
		ctrl := controller.New(db)
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
		ctrl := controller.New(db)
		repo := ctx.Params().Get("repository")
		repository, err := ctrl.Get(repo)

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

		log.Info("repositories returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": repository,
		})
	})
}
