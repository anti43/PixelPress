package controller

import (
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AdminController struct {
	system.DefaultBaseController
}

func dashboard(context *gin.Context) {
	context.HTML(http.StatusOK, "layout/admin.html", gin.H{
		"title": "Posts",
	})
}

func (t AdminController) GetAction(name string) *func(*gin.Context) {
	log.Println(name)
	switch name {
	case "dashboard":
		v := dashboard
		return &v
	default:
		return t.DefaultBaseController.GetAction(name)
	}
}
