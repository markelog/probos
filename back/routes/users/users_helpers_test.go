package users_test

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Raw("TRUNCATE users CASCADE;").Row()
}
