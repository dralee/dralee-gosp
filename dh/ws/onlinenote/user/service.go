/*
用户服务
2025.4.23 by dralee
*/
package user

import (
	"draleeonlinenote/basic"
	"draleeonlinenote/utils"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type UserService interface {
	GetAllUsers() ([]*UserImpl, error)
	GetUser(username string) (*UserImpl, error)
	GetUserById(id uint32) (*UserImpl, error)
	SaveUser(user *UserImpl) error
	UpdateUser(user *UserImpl) error
	DeleteUser(id uint32) error
	Login(w http.ResponseWriter, r *http.Request) (string, error)
	Validate(token string) error
	TokenInfo(w http.ResponseWriter, r *http.Request) (userId uint32, userName string, error error)
	Logout(w http.ResponseWriter, r *http.Request)
}

type DefaultUserService struct {
	repo UserRepository
}

func NewDefaultUserService(repo UserRepository) *DefaultUserService {
	return &DefaultUserService{repo: repo}
}

func (r *DefaultUserService) GetAllUsers() ([]*UserImpl, error) {
	return r.repo.GetAllUsers()
}

func (r *DefaultUserService) GetUser(username string) (*UserImpl, error) {
	return r.repo.GetUser(username)
}

func (r *DefaultUserService) GetUserById(id uint32) (*UserImpl, error) {
	return r.repo.GetUserById(id)
}

func (r *DefaultUserService) SaveUser(user *UserImpl) error {
	return r.repo.SaveUser(user)
}

func (r *DefaultUserService) UpdateUser(user *UserImpl) error {
	return r.repo.UpdateUser(user)
}

func (r *DefaultUserService) DeleteUser(id uint32) error {
	return r.repo.DeleteUser(id)
}

func (u *DefaultUserService) Login(w http.ResponseWriter, r *http.Request) (string, error) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		return "", fmt.Errorf("username or password is empty")
	}
	user, err := u.GetUser(username)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", fmt.Errorf("wrong password")
	}

	str := fmt.Sprintf("%d:%s", user.Id, user.UserName)
	token, err := utils.AesEncrypt([]byte(str), []byte(utils.EncryptKey))
	if err != nil {
		return "", err
	}

	var isAdmin bool = user.IsAdmin == true
	basic.Context().Login(user)
	result := basic.NewResult(0, "success", LoginResponse{Token: token, UserId: user.Id, UserName: user.UserName, IsAdmin: isAdmin})
	u.saveToken(w, r, token)
	w.Write([]byte(basic.ToJson(result)))
	return token, nil
}

func (u *DefaultUserService) saveToken(w http.ResponseWriter, r *http.Request, data string) {
	cookie := &http.Cookie{
		Name:     basic.TokenKey,
		Value:    data,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7天
	}
	http.SetCookie(w, cookie)
}

func (u *DefaultUserService) Logout(w http.ResponseWriter, r *http.Request) {
	id, _, err := u.TokenInfo(w, r)
	if err != nil {
		return
	}
	basic.Context().Logout(id)
	result := basic.NewResult(0, "success", nil)
	w.Write([]byte(basic.ToJson(result)))
}

func (u *DefaultUserService) TokenInfo(w http.ResponseWriter, r *http.Request) (userId uint32, userName string, error error) {
	token, err := r.Cookie(basic.TokenKey)
	if err != nil {
		return 0, "", err
	}
	str, err := utils.AesDecrypt([]byte(token.Value), []byte(utils.EncryptKey))
	if err != nil {
		return 0, "", err
	}
	items := strings.Split(string(str), ":")
	if len(items) == 0 {
		return 0, "", fmt.Errorf("invalid token")
	}
	userId = uint32(utils.ToInt(items[0]))
	userName = items[1]
	return userId, userName, nil
}

func (r *DefaultUserService) Validate(token string) error {
	_, err := utils.AesDecrypt([]byte(token), []byte(utils.EncryptKey))
	if err != nil {
		return err
	}
	return nil
}
