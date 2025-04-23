/*
测试加密
2025.4.23 by dralee
*/
package main

import (
	"draleeonlinenote/utils"
	"fmt"
)

func main() {
	key := []byte("1234567890123456")
	plainText := []byte("hello world")
	encrypted, err := utils.AesEncrypt(plainText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted:", encrypted)
	decrypted, err := utils.AesDecrypt([]byte(encrypted), key)
	fmt.Println("descrypted:", string(decrypted))
}
