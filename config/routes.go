package config

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"log"
)

var configRegistry *system.ControllerRegistry = nil

func Routes(registry *system.ControllerRegistry, route *gin.Engine) {

	configRegistry = registry
	r := route.Group("/pxp")
	{
		r.GET("/:controller/:action", handleDefaultRoute)
		r.GET("/:controller/:action/:id", handleDefaultRoute)
	}
}

func handleDefaultRoute(c *gin.Context) {
	controller := c.Param("controller")
	//action := c.Param("action")
	//id := c.Param("id")

	newControllerInstance, err := configRegistry.NewController(controller)
	if err != nil {
		log.Printf("unrecognized controller name: %s, cause: %s", controller, err)
	} else {
		log.Printf("controller name: %s, value %s", controller, newControllerInstance)
	}
	f := *newControllerInstance
	ff := *f.GetAction(controller)
	log.Println(ff)
	ff(c)
}
