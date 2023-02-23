package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/models"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	loginUser, err := models.GetAUserInfoByEmail(aLoginUser)
	if err != nil {
		fmt.Println(err)
	}
	if token != aLoginUser+loginUser.Pwd {
		c.JSON(http.StatusFailedDependency, Response{StatusCode: 1, StatusMsg: "Only user logged in can favorite!"})
	}

	nowVideo, err := GetAVideoById(videoId)
	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		nowVideo.IsFavorite = !nowVideo.IsFavorite
		if nowVideo.IsFavorite {
			nowVideo.FavoriteCount++
		} else {
			nowVideo.FavoriteCount--
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	allVideo, err := GetAllVideo()
	if err != nil {
		fmt.Println(err)
	}
	var videoList []Video
	videoList = nil
	for i := 0; i < len(allVideo); i++ {
		if allVideo[i].IsFavorite == true {
			videoList = append(videoList, *allVideo[i])
		}
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
