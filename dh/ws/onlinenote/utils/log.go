/*
日志
2025.4.22 by dralee
*/
package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct{}

func NewLogger(logName string) *Logger {
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%s.log", logName),
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     28,   // days
		Compress:   true, // 压缩旧文件
	}
	multiWriter := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(multiWriter) //logger) // 设置日志输出，设置同时控制台输出及文件输出
	return &Logger{}
}

func (l *Logger) Info(msg string, v ...any) {
	log.Printf(msg, v...)
}

func (l *Logger) Errorf(msg string, v ...any) {
	log.Fatalf(msg, v...) // v...是可变参数展开
}

func (l *Logger) Error(err error) {
	log.Fatal(err.Error())
}
