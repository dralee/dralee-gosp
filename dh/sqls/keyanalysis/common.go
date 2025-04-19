/*
公共
2025.4.19 by dralee
*/
package main

import (
	"keyanalysis/utils"
	"strings"
)

// 忽略文件
func IsIgnoreFile(srcPath string) bool {
	fileName := utils.FileName(srcPath)
	fileExt := utils.FileExt(srcPath)
	if !strings.HasPrefix(fileName, "FAT_") {
		return true
	}

	if strings.Contains(fileName, "xxljob") {
		return true
	}

	if fileExt != ".sql" {
		return true
	}

	num := utils.FindGroupNum(FileIndexReg, fileName, 1)
	return num < 3
}
