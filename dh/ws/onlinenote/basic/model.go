/*
基础模型
2025.4.23 by dralee
*/
package basic

type User interface {
	GetId() uint32
	GetUser() *BaseUser
}

type BaseUser struct {
	Id       uint32 `json:"id"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (u *BaseUser) GetId() uint32 {
	return u.Id
}

func (u *BaseUser) GetUser() *BaseUser {
	return u
}
