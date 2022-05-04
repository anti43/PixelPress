package main

import (
	"PixelPress/config"
	"PixelPress/controller"
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"log"
)

var registry = system.NewControllerRegistry()
var engine = gin.Default()

func initialize() {
	log.Println("PixelPress initialize()")
	registry.RegisterType(controller.TestController{})
	//fixme + name
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
