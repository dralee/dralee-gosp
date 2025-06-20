/*
package handlers
2025.6.17 by dralee
大文件上传, 分片上传
*/
package handlers

import (
	"dralee-easyupload/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CheckChunk 检查分片是否存在
func CheckChunk(c *gin.Context) {
	fileID := c.Query("file_id")
	index := c.Query("index") // 当前是第几个分片
	indexInt, _ := strconv.Atoi(index)
	chunkPath := filepath.Join("uploads", fileID, fmt.Sprintf("chunk_%d", indexInt))
	if utils.FileExists(chunkPath) {
		c.JSON(200, gin.H{"status": "chunk allready uploaded before"})
		return
	}

	c.JSON(404, gin.H{"status": "chunk not exists"})
}

// UploadChunk 上传分片
func UploadChunk(c *gin.Context) {
	fileID := c.PostForm("file_id")
	index := c.PostForm("index") // 当前是第几个分片
	indexInt, _ := strconv.Atoi(index)

	file, _, err := c.Request.FormFile("chunk")
	if err != nil {
		c.JSON(400, gin.H{"error": "No chunk received"})
		return
	}
	defer file.Close()

	dir := filepath.Join("uploads", fileID)
	os.MkdirAll(dir, os.ModePerm)

	// 说明是断点续传
	chunkPath := filepath.Join(dir, fmt.Sprintf("chunk_%d", indexInt))
	if utils.FileExists(chunkPath) {
		c.JSON(200, gin.H{"status": "chunk allready uploaded before"})
		return
	}

	out, err := os.Create(chunkPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot create chunk"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save chunk"})
		return
	}

	c.JSON(200, gin.H{"status": "chunk uploaded"})
}
