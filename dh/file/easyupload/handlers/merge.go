/*
package handlers
2025.6.17 by dralee
大文件上传, 合并分片
*/
package handlers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MergeChunks(c *gin.Context) {
	fileID := c.PostForm("file_id")
	totalStr := c.PostForm("total_chunks")
	filename := c.PostForm("filename")
	root := "uploads"

	totalChunks, _ := strconv.Atoi(totalStr)
	chunkDir := filepath.Join(root, fileID)
	mergedPath := filepath.Join(root, fileID+"_"+filename)
	fmt.Printf("totalChunks: %d, chunkDir: %s, mergedPath: %s\n", totalChunks, chunkDir, mergedPath)

	mergedFile, err := os.Create(mergedPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot create merged file"})
		return
	}

	defer func() {
		mergedFile.Close()
		if chunkDir != root {
			os.RemoveAll(chunkDir)
			//clearDir(chunkDir)
		}
	}()

	// if utils.FileExists(mergedPath) {
	// 	os.Remove(mergedPath)
	// }

	for i := 0; i < totalChunks; i++ {
		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("chunk_%d", i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Missing chunk %d", i)})
			return
		}
		io.Copy(mergedFile, chunkFile)
		chunkFile.Close()
	}

	c.JSON(200, gin.H{"status": "merged", "file": mergedPath})
}
