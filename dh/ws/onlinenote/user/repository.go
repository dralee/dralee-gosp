/*
用户仓储
2025.4.23 by dralee
*/
package user

import (
	"database/sql"
	"time"
)

type UserRepository interface {
	GetAllUsers() ([]*UserImpl, error)
	GetUser(UserName string) (*UserImpl, error)
	GetUserById(id uint32) (*UserImpl, error)
	SaveUser(user *UserImpl) error
	UpdateUser(user *UserImpl) error
	DeleteUser(id uint32) error
}

type DefaultUserRepository struct {
	db         *sql.DB
	connString string
}

func NewDefaultUserRepository(connString string) (*DefaultUserRepository, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	repo := DefaultUserRepository{connString: connString, db: db}
	return &repo, nil
}

func (r *DefaultUserRepository) GetAllUsers() ([]*UserImpl, error) {
	var sql = "SELECT `id`,`username`, `password`, `is_admin`, `is_enabled`, `creation_time` FROM `user`;"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*UserImpl
	for rows.Next() {
		var user UserImpl
		err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.IsAdmin, &user.IsEnabled, &user.CreationTime)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *DefaultUserRepository) GetUser(UserName string) (*UserImpl, error) {
	var sql = "SELECT `id`,`username`, `password`, `is_admin`, `is_enabled`, `creation_time` FROM `user` WHERE `UserName` = ?;"
	var user UserImpl
	err := r.db.QueryRow(sql, UserName).Scan(&user.Id, &user.UserName, &user.Password, &user.IsAdmin, &user.IsEnabled, &user.CreationTime)
	return &user, err
}

func (r *DefaultUserRepository) GetUserById(id uint32) (*UserImpl, error) {
	var sql = "SELECT `id`,`username`, `password`, `is_admin`, `is_enabled`, `creation_time` FROM `user` WHERE `id` = ?;"
	var user UserImpl
	err := r.db.QueryRow(sql, id).Scan(&user.Id, &user.UserName, &user.Password, &user.IsAdmin, &user.IsEnabled, &user.CreationTime)
	return &user, err
}

func (r *DefaultUserRepository) SaveUser(user *UserImpl) error {
	var sql = "INSERT INTO `user` (`username`, `password`, `is_admin`, `is_enabled`, `creation_time`) VALUES (?, ?, ?, ?, ?);"
	_, err := r.db.Exec(sql, user.UserName, user.Password, user.IsAdmin, user.IsEnabled, user.CreationTime)
	return err
}

func (r *DefaultUserRepository) UpdateUser(user *UserImpl) error {
	var sql = "UPDATE `user` SET `username` = ?, `password` = ?, `is_admin` = ?, `is_enabled` = ?, last_modification_time = ? WHERE `UserName` = ?;"
	_, err := r.db.Exec(sql, user.UserName, user.Password, user.IsAdmin, user.IsEnabled, time.Now().Unix(), user.UserName)
	return err
}

func (r *DefaultUserRepository) DeleteUser(id uint32) error {
	var sql = "DELETE FROM `user` WHERE `id` = ?;"
	_, err := r.db.Exec(sql, id)
	return err
}
