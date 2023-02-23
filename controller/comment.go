package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/models"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId := c.Query("video_id")
	logUser, err := models.GetAUserInfoByEmail(aLoginUser)
	if err != nil {
		fmt.Println(err)
	}
	if token == aLoginUser+logUser.Pwd {
		if actionType == "1" {
			text := c.Query("comment_text")
			vdo, err := GetAVideoById(videoId)
			if err != nil {
				fmt.Println(err)
			}
			cmts, err := GetAllComment()
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         int64(len(cmts)) + 1,
					User:       vdo.Author,
					Content:    text,
					CreateDate: time.Time{}.String(),
				}})
			return
		}
		cid := c.Query("comment_id")
		err = DeleteAComment(cid)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	vdoId := c.Query("video_id")
	allComments, err := GetAllComment()
	var commentList []Comment
	commentList = nil
	for i := 0; i < len(allComments); i++ {
		if vdoId == strconv.FormatInt(allComments[i].VideoId, 10) {
			commentList = append(commentList, *allComments[i])
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
