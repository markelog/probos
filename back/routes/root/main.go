package root

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/probos/back/database/models"
	"github.com/sirupsen/logrus"
)

var Repository models.Repository

// Up root route
func Up(app *iris.Application, db *gorm.DB, logger *logrus.Logger) {
	app.Get("/", func(ctx iris.Context) {
		db.First(&Repository)
		ctx.HTML("<h1>" + Repository.Name + "</h1>")
	})
}
