package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController interface {
}

type BaseControllerWithActions interface {
	GetAction(name string) *func(*gin.Context)
}

type DefaultBaseController struct {
	BaseController
	BaseControllerWithActions
}

func base(context *gin.Context) {
	context.String(http.StatusOK, "base")
}

func (DefaultBaseController) GetAction(name string) *func(*gin.Context) {
	switch name {
	default:
		v := base
		return &v
	}
}
