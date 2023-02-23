package controller

import (
	"bytes"
	"io"
	"net/http"
	"simple-demo/dao"
	"simple-demo/models"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64      `json:"id,omitempty"`
	Author        User       `json:"author"`
	PlayUrl       string     `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	IsFavorite    bool       `json:"is_favorite,omitempty"`
	Comments      []*Comment `json:"-"`               // 评论信息
	CreatedAt     time.Time  `json:"-"`               // 创建时间
	UpdatedAt     time.Time  `json:"-"`               // 更新时间
	Title         string     `json:"title,omitempty"` // 视频标题
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"type:bigint;not null"`
	User       User   `json:"cuser" gorm:"type:bigint;not null"`
	Content    string `json:"content,omitempty" gorm:"type:varchar(255);not null"`
	CreateDate string `json:"create_date,omitempty" gorm:"type:varchar(32);not null"`
	VideoId    int64  `json:"video_id" gorm:"type:bigint;not null"`
}

func CreateAComment(comment *Comment) (err error) {
	err = dao.DB.Create(&comment).Error
	return
}

func GetAllComment() (commentList []*Comment, err error) {
	if err = dao.DB.Find(&commentList).Error; err != nil {
		return nil, err
	}
	return
}

func GetACommentById(id string) (comment *Comment, err error) {
	comment = new(Comment)
	if err = dao.DB.Debug().Where("id=?", id).First(comment).Error; err != nil {
		return nil, err
	}
	return
}

func GetACommentByName(id string) (comment *Comment, err error) {
	comment = new(Comment)
	if err = dao.DB.Debug().Where("user=?", id).First(comment).Error; err != nil {
		return nil, err
	}
	return
}

func GetACommentByVideoId(id string) (comment *Comment, err error) {
	comment = new(Comment)
	if err = dao.DB.Debug().Where("video_id=?", id).First(comment).Error; err != nil {
		return nil, err
	}
	return
}
func UpdateAComment(comment *Comment) (err error) {
	err = dao.DB.Save(comment).Error
	return
}

func DeleteAComment(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Comment{}).Error
	return
}

type FeedVideoList struct {
	Videos   []*Video `json:"video_list,omitempty"`
	NextTime int64    `json:"next_time,omitempty"`
}
type QueryFeedVideoListFlow struct {
	userId     int64
	latestTime time.Time
	videos     []*Video
	nextTime   int64
	feedVideo  *FeedVideoList
}
type User struct {
	Id            int64  `json:"id,omitempty" gorm:"type:bigint;not null"`
	Name          string `json:"name,omitempty" gorm:"type:varchar(32);not null"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"type:bigint;not null"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"type:bigint;not null"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"type:tinyint(1);not null"`
	Signature     string `json:"signature" gorm:"type:varchar(32);"`
	TotalFavorite int64  `json:"total_favorite" gorm:"type:bigint;"`
	WorkCount     int64  `json:"work_count" gorm:"type:bigint;"`
	FavoriteCount int64  `json:"favorite_count" gorm:"type:bigint;"`
}

func CreateAUser(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

func GetAllUser() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUserByName(name string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("name=?", name).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUserById(id string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUser(user *User) (err error) {
	err = dao.DB.Save(user).Error
	return
}

func DeleteAUser(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}

func CreateAVideo(video *Video) (err error) {
	err = dao.DB.Create(&video).Error
	return
}

func GetAllVideo() (videoList []*Video, err error) {
	if err = dao.DB.Find(&videoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAVideoByName(name string) (video *Video, err error) {
	video = new(Video)
	if err = dao.DB.Debug().Where("name=?", name).First(video).Error; err != nil {
		return nil, err
	}
	return
}

func GetAVideoById(id string) (video *Video, err error) {
	video = new(Video)
	if err = dao.DB.Debug().Where("id=?", id).First(video).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAVideo(video *Video) (err error) {
	err = dao.DB.Save(video).Error
	return
}

func DeleteAVideo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Video{}).Error
	return
}

func DownLoadAsBytes(url string) (video []byte, err error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//// Create output file
	//out, err := os.Create(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer out.Close()
	// copy stream
	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, resp.Body)
	video = buff.Bytes()
	if err != nil {
		panic(err)
	}
	return
}

func ConvertVideo2VideoInfo(video *Video) (videoInfo models.VideoInfo, err error) {
	vdo, err := DownLoadAsBytes(video.PlayUrl)
	videoInfo = models.VideoInfo{
		Id:    video.Id,
		Vname: video.Author.Name + time.Time{}.String(),
		Video: vdo,
	}
	return
}

func ConvertVideoInfo2Video(videoInfo *models.VideoInfo) (video Video, err error) {
	//vdo, err := DownLoadAsBytes(video.PlayUrl)
	author, err := GetAUserById(videoInfo.Author)
	video = Video{
		Id:            videoInfo.Id,
		Author:        *author,
		PlayUrl:       "192.168.2.169",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	return
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
