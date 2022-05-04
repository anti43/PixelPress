package controller

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
	system.DefaultBaseController
}

func test(context *gin.Context) {
	context.String(http.StatusOK, "test")
}

//fixme do via reflection
func (t TestController) GetAction(name string) *func(*gin.Context) {
	switch name {
	case "test":
		v := test
		return &v
	default:
		return t.DefaultBaseController.GetAction(name)
	}
}
