package models

import (
	"simple-demo/dao"
)

// User Model
type User struct {
	id               int64  `json:"id"`
	follow_count     int64  `json:"follow_count"`
	follower_count   int64  `json:"follower_count"`
	is_follow        bool   `json:"is_follow"`
	avatar           string `json:"avatar"`
	background_image string `json:"background_image"`
	signature        string `json:"signature"`
	total_favorited  int64  `json:"total_favorited"`
	work_count       int64  `json:"work_count"`
	favorate_count   int64  `json:"favorate_count""`
	email            string `json:"email"`
	password         string `json:"password"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

func GetAllTodo() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(user *User) (err error) {
	err = dao.DB.Save(user).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
