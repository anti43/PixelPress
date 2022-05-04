package main

import (
	"PixelPress/config"
	"PixelPress/controller"
	"PixelPress/system"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var registry = system.NewControllerRegistry()
var engine = gin.Default()
var dbConnection *gorm.DB

func initialize() {
	engine.SetTrustedProxies(nil)
	engine.HTMLRender = gintemplate.Default()

	log.Println("PixelPress initialize()")
	registry.RegisterNamedType("controller.BaseController", system.DefaultBaseController{})
	registry.RegisterType(controller.TestController{})
	registry.RegisterType(controller.AdminController{})

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbConnection = db
}

func main() {
	initialize()

	config.Routes(registry, engine, dbConnection)
	err := engine.Run()
	if err != nil {
		panic(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
