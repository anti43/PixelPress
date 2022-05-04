package config

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
)

var configRegistry *system.ControllerRegistry = nil

func handleDefaultRoute(c *gin.Context) {
	system.HandleDefaultRoute(c, configRegistry)
}

func Routes(registry *system.ControllerRegistry, route *gin.Engine) {
	configRegistry = registry

	r := route.Group("/pxp")
	{
		r.GET("/:controller/:action", handleDefaultRoute)
		r.GET("/:controller/:action/:id", handleDefaultRoute)
	}
}

