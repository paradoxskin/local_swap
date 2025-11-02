package controller

import "github.com/gin-gonic/gin"
import "local_swap/service"
import "fmt"
import "time"
import "strings"
import "sync"

var mu sync.Mutex
var count uint8

func GetPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func PostFileList(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"ip": service.GetIP(),
		"files": service.GetFiles(),
		"texts": service.GetText(),
	})
}

func UploadFile(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, "error")
		return
	}
	ls := strings.Split(file.Filename, ".")
	var filename string
	if len(ls) != 1 {
		lst := ls[len(ls) - 1]
		filename = fmt.Sprintf("%x_%02x.%s", time.Now().Unix(), count, lst)
	} else {
		filename = fmt.Sprintf("%x_%02x", count, time.Now().Unix())
	}
	count += 1
	c.SaveUploadedFile(file, fmt.Sprintf("./static/files/%s", filename))
	c.JSON(200, map[string]string{
		"filename": filename,
	})
}

func PostText(c *gin.Context) {
	text := c.PostForm("text")
	service.WriteText(text)
}

func CleanAll(c *gin.Context) {
	service.CleanAll()
}
