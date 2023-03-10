package controller

import "github.com/gin-gonic/gin"
import "local_swap/service"

func GetPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func PostFileList(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"ip": service.GetIP(),
	})
}

func UploadFile(c *gin.Context) {
	
}

func Download(c *gin.Context) {
	
}
