/*
用户实体
2025.4.23 by dralee
*/
package user

import (
	"draleeonlinenote/basic"
)

type UserImpl struct {
	basic.BaseUser
	Password     string
	IsAdmin      basic.BoolType
	IsEnabled    basic.BoolType
	CreationTime uint64
}

func (u *UserImpl) GetId() uint32 {
	return u.Id
}

func (u *UserImpl) GetUser() *basic.BaseUser {
	return &u.BaseUser
}
