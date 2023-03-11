package router

import (
	"github.com/gin-gonic/gin"
	"local_swap/controller"
)

func Build() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templ/*")

	r.Static("/js", "static/js")
	r.Static("/css", "static/css")
	r.Static("/font", "static/fonts")
	r.Static("/files", "static/files")
	r.Static("/svg", "static/svg")

	r.GET("/", controller.GetPage)
	r.POST("/", controller.PostFileList)
	r.POST("/text", controller.PostText)
	r.POST("/upload", controller.UploadFile)
	r.POST("/download", controller.Download)
	return r
}
