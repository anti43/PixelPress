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

//fixme how to avoid this boilerplate code
func (TestController) getAction(name string) *func(*gin.Context) {
	switch name {
	case "test":
		v := test
		return &v
	default:
		return nil
	}
}
