/*
package handlers
2025.6.17 by dralee
大文件上传, 首页
*/
package handlers

import "github.com/gin-gonic/gin"

func HomePage(c *gin.Context) {
	c.HTML(200, "web/index.html", nil)
}
