package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	controller "github.com/markelog/probos/back/controllers/repository"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

type postRepository struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

var schema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"repository": {"type": "string", "minLength": 1}
	},
	"required": ["name", "repository"]
}`)

func validate(params *postRepository) (*gojsonschema.Result, *iris.Map) {
	var (
		paramsLoader = gojsonschema.NewGoLoader(params)
		check, _     = gojsonschema.Validate(schema, paramsLoader)

		errors  []string
		payload *iris.Map
	)

	if check.Valid() == false {
		for _, desc := range check.Errors() {
			errors = append(errors, desc.String())
		}

		payload = &iris.Map{"errors": errors}

		return check, payload
	}

	return check, nil
}

// Up Repository route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/repositories", func(ctx iris.Context) {
		var params postRepository
		ctx.ReadJSON(&params)

		validation, errors := validate(&params)

		if validation.Valid() == false {
			log.WithFields(logrus.Fields{
				"Repository": params.Name,
				"repository": params.Repository,
			}).Error("Params are not valid")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Params are not valid",
				"payload": errors,
			})

			return
		}

		ctrl := controller.New(db)
		result, err := ctrl.Create(params.Name, params.Repository)

		if err != nil {
			log.WithFields(logrus.Fields{
				"Repository": params.Name,
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
		Repositories, err := ctrl.List()

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

		if len(Repositories) == 0 {
			log.Error("Can't find any Repositories")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": iris.Map{},
			})

			return
		}

		log.Info("Repositories returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": Repositories,
		})

		return
	})
}
