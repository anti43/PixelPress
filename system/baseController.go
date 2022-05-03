package system

import "github.com/gin-gonic/gin"

type BaseController interface {
}

type BaseControllerWithActions interface {
	GetAction(name string) *func(*gin.Context)
}

type DefaultBaseController struct {
	BaseController
	BaseControllerWithActions
}

func (*DefaultBaseController) GetAction(name string) *func(*gin.Context) {
	switch name {
	default:
		return nil
	}
}
