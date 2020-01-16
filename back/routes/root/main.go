package root

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/markelog/probos/back/database/models"
	"github.com/sirupsen/logrus"
)

var repository models.Repository

// Up root route
func Up(app *iris.Application, db *gorm.DB, logger *logrus.Logger) {
	app.Get("/", func(ctx iris.Context) {
		db.First(&repository)
		ctx.HTML("<h1>" + repository.Name + "</h1>")
	})
}
