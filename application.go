package main

import (
	"PixelPress/config"
	"PixelPress/controller"
	"PixelPress/system"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"log"
)

var registry = system.NewControllerRegistry()
var engine = gin.Default()

func initialize() {
	engine.SetTrustedProxies(nil)
	engine.HTMLRender = gintemplate.Default()
	
	log.Println("PixelPress initialize()")
	registry.RegisterNamedType("controller.BaseController", system.DefaultBaseController{})
	registry.RegisterType(controller.TestController{})
	registry.RegisterType(controller.AdminController{})

}

func main() {
	initialize()

	config.Routes(registry, engine)
	err := engine.Run()
	if err != nil {
		panic(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
