package models

import (
	"simple-demo/dao"
)

// UserInfo Model
type UserInfo struct {
	//gorm.Model
	Id    int64  `json:"id" gorm:"type:bigint;not null"`
	Email string `json:"email" gorm:"type:varchar(32);not null"`
	//follow_count     int64  `json:"follow_count"`
	//follower_count   int64  `json:"follower_count"`
	//is_follow        bool   `json:"is_follow"`
	//avatar           string `json:"avatar"`
	//background_image string `json:"background_image"`
	//signature        string `json:"signature"`
	//total_favorited  int64  `json:"total_favorited"`
	//work_count       int64  `json:"work_count"`
	//favorate_count   int64  `json:"favorate_count""`
	//email            string `json:"email"`
	Pwd string `json:"pwd" gorm:"type:varchar(32);not null"`
}

/*
	User这个Model的增删改查操作都放在这里
*/
// CreateAUser 创建user
func CreateAUserInfo(userInfo *UserInfo) (err error) {
	err = dao.DB.Create(&userInfo).Error
	return
}

func GetAllUserInfo() (userInfoList []*UserInfo, err error) {
	if err = dao.DB.Find(&userInfoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUserInfoByEmail(email string) (userInfo *UserInfo, err error) {
	userInfo = new(UserInfo)
	if err = dao.DB.Debug().Where("email=?", email).First(userInfo).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUserInfoById(id string) (userInfo *UserInfo, err error) {
	userInfo = new(UserInfo)
	if err = dao.DB.Debug().Where("id=?", id).First(userInfo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUserInfo(userInfo *UserInfo) (err error) {
	err = dao.DB.Save(userInfo).Error
	return
}

func DeleteAUserInfo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&UserInfo{}).Error
	return
}
