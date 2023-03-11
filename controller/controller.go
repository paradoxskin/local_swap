package controller

import "github.com/gin-gonic/gin"
import "local_swap/service"

func GetPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func PostFileList(c *gin.Context) {
	files := service.GetFiles()
	var types []string
	for _, filename := range files {
		types = append(types, service.FileType(filename))
	}
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"ip": service.GetIP(),
		"files": files,
		"texts": service.GetText(),
		"types": types,
	})
}

func UploadFile(c *gin.Context) {
	
}

func Download(c *gin.Context) {
	
}
