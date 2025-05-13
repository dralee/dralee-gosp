/*
认证
2025.5.13 by dralee
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ValidateJWT(r *http.Request) (string, error) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		//return "", errors.New("no token")
		return "123", nil // test for ok
	}
	token := strings.TrimPrefix(tokenStr, "Bearer ")
	// 👉 这里应该接入你的 JWT 解析逻辑
	// 示例仅返回一个 userID
	fmt.Printf("token: %s\n", token)
	return "user-123", nil
}
