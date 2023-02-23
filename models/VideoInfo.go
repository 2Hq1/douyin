package models

import (
	"simple-demo/dao"
)

// UserInfo Model
type VideoInfo struct {
	//gorm.Model
	Id     int64  `json:"id" gorm:"type:bigint;not null"`
	Vname  string `json:"name" gorm:"type:varchar(32);not null"`
	Video  []byte `json:"video" gorm:"type:longblob;not null"`
	Author string `json:"name" gorm:"type:varchar(32);not null"`
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
	//Pwd string `json:"pwd" gorm:"type:varchar(32);not null"`
}

/*
	User这个Model的增删改查操作都放在这里
*/
// CreateAUser 创建user
func CreateAVideoInfo(videoInfo *VideoInfo) (err error) {
	err = dao.DB.Create(&videoInfo).Error
	return
}

func GetAllVideoInfo() (videoInfoList []*VideoInfo, err error) {
	if err = dao.DB.Find(&videoInfoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAuthorVideoInfo(name string) (videoInfoList []*VideoInfo, err error) {
	if err = dao.DB.Find(&videoInfoList).Where("author-?", name).Error; err != nil {
		return nil, err
	}
	return
}

func GetAVideoInfoByName(name string) (videoInfo *VideoInfo, err error) {
	videoInfo = new(VideoInfo)
	if err = dao.DB.Debug().Where("name=?", name).First(videoInfo).Error; err != nil {
		return nil, err
	}
	return
}

func GetAVideoInfoById(id string) (videoInfo *VideoInfo, err error) {
	videoInfo = new(VideoInfo)
	if err = dao.DB.Debug().Where("id=?", id).First(videoInfo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAVideoInfo(videoInfo *VideoInfo) (err error) {
	err = dao.DB.Save(videoInfo).Error
	return
}

func DeleteAVideoInfo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&VideoInfo{}).Error
	return
}
