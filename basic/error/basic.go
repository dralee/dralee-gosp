/*
错误处理：

	error接口：标准错误表示（任何实现了error接口的Error方法的类型都可作为错误）
	显式返回值：通过函数的返回值返回错误
	自定义错误：可通过标准库或自定义方式创建错误
	panic和recover：处理不可恢复的严格错误
		panic：主动抛出错误
		recover：捕获错误

fmt错误格式化输出：

	%v: 默认格式
	%+v：如果支持，显式详细的错误信息
	%s：作为字符串输出

2025.4.17 by dralee
*/
package main

import (
	"errors"
	"fmt"
)

func devide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

type DevideError struct {
	dividend int
	divisor  int
}

func (e *DevideError) Error() string {
	return fmt.Sprintf("cannot devide %d by %d", e.dividend, e.divisor)
}

func devide1(x, y int) (int, error) {
	if y == 0 {
		return 0, &DevideError{x, y}
	}
	return x / y, nil
}

type MyDevideError struct {
	dividee int
	divider int
}

func (e *MyDevideError) Error() string {
	sf := `
	Cannot proceed, the devider is zero.
	dividee: %d
	divider: %d
	`
	return fmt.Sprintf(sf, e.dividee, e.divider)
}

func devide2(x, y int) (result int, errMsg string) {
	if y == 0 {
		d := MyDevideError{x, y}
		return 0, d.Error()
	}
	return x / y, ""
}

var ErrorNotFound = errors.New("not found")

func find(id int) error {
	return fmt.Errorf("db error: id %d not found, %w", id, ErrorNotFound)
}

type MyError struct {
	code int
	msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.code, e.msg)
}
func getError() error {
	return &MyError{
		code: 404,
		msg:  "Not Found",
	}
}

func safeErr() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	panic("a problem occured")
}

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)

	r, err := devide(1, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", r)
	}

	_, err = devide1(1, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", r)
	}

	_, errMsg := devide2(1, 0)
	if errMsg != "" {
		fmt.Println(errMsg)
	} else {
		fmt.Println("result:", r)
	}

	err = find(1)
	if errors.Is(err, ErrorNotFound) { // 是否特定错误或由该错误包装而成
		fmt.Println("data not found")
	} else {
		fmt.Println("unknown error", err)
	}

	err = getError()
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("my error - code: %d, msg: %s\n", myErr.code, myErr.msg)
	}

	safeErr()
}
