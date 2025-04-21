/*
公共
2025.4.19 by dralee
*/
package main

import (
	"errors"
	"keyanalysis/utils"
	"strings"
)

type KeyOperateType int

const (
	Scan KeyOperateType = iota
	Replace
)

var keyOperateTypeMap = map[string]KeyOperateType{
	"scan":    Scan,
	"replace": Replace,
}

func ParseKeyOperateType(s string) (KeyOperateType, error) {
	key := strings.ToLower(s)
	if keyOperateType, ok := keyOperateTypeMap[key]; ok {
		return keyOperateType, nil
	}
	return 0, errors.New("invalid key operate type")
}

const (
	OutputPath = "./data"
	SrcKeyFile = "./data/src-key.txt"
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
	// 由于基础数据太多，暂时忽略
	/*if strings.Contains(fileName, "_bds_") {
		return true
	}*/

	if fileExt != ".sql" {
		return true
	}

	num := utils.FindGroupNum(FileIndexReg, fileName, 1)
	return num < 3
}

var logger *utils.Logger

// 初始化日志
func initLogger() {
	logger = utils.NewLogger("keyanalysis")
}

func Info(msg string, v ...any) {
	logger.Info(msg, v...)
}
func Errorf(msg string, v ...any) {
	logger.Errorf(msg, v...)
}
func Error(err error) {
	logger.Error(err)
}
