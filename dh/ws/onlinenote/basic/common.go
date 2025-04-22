/*
公共部分
2025.4.22 by dralee
*/
package onlinenote

import (
	"draleeonlinenote/utils"
)

var (
	Newline = []byte{'\n'}
	Space   = []byte{' '}
)

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
