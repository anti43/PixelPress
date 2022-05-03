package config

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"log"
)

var configRegistry *system.Registry = nil

func Routes(registry *system.Registry, route *gin.Engine) {

	configRegistry = registry
	r := route.Group("/pxp")
	{
		r.GET("/:controller/:action", handleDefaultRoute)
		r.GET("/:controller/:action/:id", handleDefaultRoute)
	}
}

func handleDefaultRoute(c *gin.Context) {
	controller := c.Param("controller")
	action := c.Param("action")
	//id := c.Param("id")

	controllerPointer, err := configRegistry.NewController(controller)
	if err != nil {
		log.Print("unrecognized controller name: %s, cause: %s", controller, err)
	}

	defbc := *controllerPointer
	var x = defbc.(system.BaseControllerWithActions)
	f := *x.GetAction(action)
	f(c)
}
