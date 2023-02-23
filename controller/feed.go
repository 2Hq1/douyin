package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/models"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

var (
	feedCount int64
)

const (
	MaxVideoNum = 30
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videoPtrList, errGetVideoList := GetAllVideo()
	videoInfoList, err := models.GetAllVideoInfo()
	if err != nil {
		fmt.Println(err)
	}
	if errGetVideoList != nil {
		fmt.Println(errGetVideoList)
	}
	//cfg := setting.Conf.FeedConfig
	var videoList []Video
	videoList = nil
	if videoPtrList == nil {
		videoList = DemoVideos
	} else {
		var vdo Video
		//dsn := fmt.Sprintf("http://%s:8080/static/", cfg.Ip)
		for i := 0; i < len(videoInfoList); i++ {
			vdo = *videoPtrList[i]
			vdo.PlayUrl = "http://192.168.43.89:8080/static/" + videoInfoList[i].Vname
			videoList = append(videoList, vdo)
		}
	}
	//fmt.Printf("videoList:%#v\n", videoList)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
