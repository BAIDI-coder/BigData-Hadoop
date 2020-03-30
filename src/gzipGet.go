package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := ioutil.ReadFile("./download/" + filename)
		if err != nil {
			c.JSON(404, gin.H{
				"msg": "can't find source, please call William 1105282405@qq.com",
			})
		}
		c.Data(200, "application/x-gzip", file)
	})
	router.Run(":80")
}
