package system

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleDefaultRoute(c *gin.Context, configRegistry *ControllerRegistry) {
	controller := c.Param("controller")
	action := c.Param("action")
	//id := c.Param("id")

	newControllerInstance, err := configRegistry.NewController(controller)
	if err != nil {
		log.Printf("unrecognized controller name: %s, cause: %s", controller, err)
	} else {
		log.Printf("controller name: %s, value %s", controller, newControllerInstance)
	}
	f := *newControllerInstance
	ff := *f.GetAction(action)
	ff(c)
}
