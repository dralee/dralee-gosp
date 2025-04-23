/*
公共部分
2025.4.22 by dralee
*/
package basic

import (
	"draleeonlinenote/utils"
	"encoding/json"
)

const (
	TokenKey = "online.tk"
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

type Result interface {
	GetCode() int
	GetMsg() string
}

type BaseResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataResult struct {
	BaseResult
	Data any `json:"data"`
}

func (r *BaseResult) GetCode() int {
	return r.Code
}
func (r *BaseResult) GetMsg() string {
	return r.Msg
}
func (r *DataResult) GetCode() int {
	return r.Code
}
func (r *DataResult) GetMsg() string {
	return r.Msg
}

func NewResult(code int, msg string, data any) Result {
	r := BaseResult{
		Code: code,
		Msg:  msg,
	}
	if data == nil {
		return &r
	}
	return &DataResult{
		BaseResult: r,
		Data:       data,
	}
}

func Success(data any) Result {
	return NewResult(0, "success", data)
}

func SuccessMsg(msg string) Result {
	return NewResult(0, msg, nil)
}

func Fail(code int, msg string) Result {
	return NewResult(code, msg, nil)
}

func ToJson(result any) []byte {
	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return data
}

func FromJson(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
