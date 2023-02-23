package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"simple-demo/models"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Println(token)
	fmt.Printf("token:%#v\n", token)
	//publish.PublishAction
	loginUser, exist := usersLoginInfo[token]
	if !exist {
		fmt.Println(exist)
	}
	getUser, errGetUser := GetAUserById(strconv.FormatInt(loginUser.Id, 10))
	if errGetUser != nil {
		fmt.Printf("user:%#v\n", getUser)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	errSave := c.SaveUploadedFile(data, saveFile)
	VideoList, errGetVideoList := models.GetAllVideoInfo()
	fmt.Printf("videoList:%#v\n", VideoList)
	if errGetVideoList != nil {
		fmt.Println(errGetVideoList)
	}
	newId := len(VideoList) + 1
	file, errOpen := data.Open()
	if errOpen != nil {
		fmt.Println(errOpen)
	}
	buff := new(bytes.Buffer)
	_, _ = io.Copy(buff, file)
	video := buff.Bytes()
	newVideo := models.VideoInfo{
		Id:     int64(newId),
		Vname:  finalName,
		Video:  video,
		Author: aLoginUser,
	}
	tmpUser, err := GetAUserByName(aLoginUser)
	newVdo := Video{
		Id:            int64(newId),
		PlayUrl:       "http://192.168.2.169:8080/static/" + newVideo.Vname,
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Author:        *tmpUser,
	}
	errCreate := models.CreateAVideoInfo(&newVideo)
	err = CreateAVideo(&newVdo)
	if err != nil {
		fmt.Println(err)
	}
	if errSave != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	if errCreate != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	authorVideoInfo, err := models.GetAuthorVideoInfo(aLoginUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	allVideo, err := GetAllVideo()
	var videoList []Video
	videoList = nil
	for i := 0; i < len(allVideo); i++ {
		for j := 0; j < len(authorVideoInfo); j++ {
			if allVideo[i].Id == authorVideoInfo[j].Id {
				videoList = append(videoList, *allVideo[i])
			}
		}
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
