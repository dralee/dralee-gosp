/*
package main
2025.6.17 by dralee
大文件上传, 入口
*/
package main

import (
	"dralee-easyupload/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 绑定 HTML 模板目录
	r.LoadHTMLGlob("web/page/*")

	// 静态文件，如 JS/CSS
	r.Static("static", "web/static")

	// 显示上传页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "分片上传大文件",
		})
	})

	r.GET("/checkchunk", handlers.CheckChunk)
	r.POST("/upload", handlers.UploadChunk)
	r.POST("/merge", handlers.MergeChunks)

	r.Run(":8080")
}
