/*
传输对象
2025.4.23 by dralee
*/
package user

type LoginResponse struct {
	Token    string `json:"token"`
	UserId   uint32 `json:"userId"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
}
