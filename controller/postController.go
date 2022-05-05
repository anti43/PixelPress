package controller

import (
	"PixelPress/model"
	"PixelPress/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strings"
)

type PostController struct {
	system.DefaultBaseController
}

func get(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)
	idParam := context.Param("id")
	post := model.Post{}
	dbConnection.First(&post, idParam)
	context.HTML(http.StatusOK, "page/post.html", gin.H{
		"title": strings.Title(reflect.TypeOf(model.Post{}).String()),
		"data":  &post,
	})
}

func deletePost(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)
	idParam := context.Param("id")
	post := model.Post{}
	dbConnection.First(&post, idParam)
	dbConnection.Delete(&post)

	list(context)
}

func list(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)
	var posts []model.Post
	dbConnection.Find(&posts)
	context.HTML(http.StatusOK, "layout/admin.html", gin.H{
		"title": strings.Title(reflect.TypeOf(model.Post{}).String()),
		"data":  posts,
	})
}

func update(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)
	idParam := context.Param("id")

	post := model.Post{}
	dbConnection.First(&post, idParam)
	post.Description = context.Param("post.Description")
	post.Text = context.Param("post.Text")
	dbConnection.Save(&post)

	context.HTML(http.StatusOK, "layout/admin.html", gin.H{
		"title": post.Description,
		"data":  post,
	})
}

func create(context *gin.Context) {
	dbConnection := context.MustGet("dbConnection").(*gorm.DB)
	post := model.Post{}
	post.Description = context.Param("post.Description")
	post.Text = context.Param("post.Text")

	dbConnection.Create(&post)

	context.HTML(http.StatusOK, "layout/admin.html", gin.H{
		"title": post.Description,
		"data":  post,
	})
}

func (t PostController) GetAction(name string) *func(*gin.Context) {

	switch name {
	case "get":
		v := get
		return &v
	case "update":
		v := update
		return &v
	case "delete":
		v := deletePost
		return &v
	case "list":
		v := list
		return &v
	case "create":
		v := create
		return &v
	default:
		return t.DefaultBaseController.GetAction(name)
	}
}
