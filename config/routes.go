package config

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var configRegistry *system.ControllerRegistry = nil
var dbConnection *gorm.DB

func handleDefaultRoute(c *gin.Context) {
	c.Set("dbConnection", dbConnection)
	system.HandleDefaultRoute(c, configRegistry)
}

func Routes(registry *system.ControllerRegistry, route *gin.Engine, connection *gorm.DB) {
	configRegistry = registry
	dbConnection = connection

	r := route.Group("/pxp")
	{
		r.GET("/:controller/:action", handleDefaultRoute)
		r.GET("/:controller/:action/:id", handleDefaultRoute)
	}
}
