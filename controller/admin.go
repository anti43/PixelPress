package controller

import (
	"PixelPress/model"
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AdminController struct {
	system.DefaultBaseController
}

func dashboard(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)

	context.HTML(http.StatusOK, "layout/admin.html", gin.H{
		"title": "Posts",
		"posts": dbConnection.Find(&model.Post{}),
	})
}

func (t AdminController) GetAction(name string) *func(*gin.Context) {

	switch name {
	case "dashboard":
		v := dashboard
		return &v
	default:
		return t.DefaultBaseController.GetAction(name)
	}
}
