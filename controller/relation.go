package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/dao"
	"simple-demo/models"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type Relation struct {
	IdFrom int64 `json:"id_from,omitempty" gorm:"type:bigint;not null"`
	IdTo   int64 `json:"id_to,omitempty" gorm:"type:bigint;not null"`
}

func CreateRalation(relation *Relation) (err error) {
	err = dao.DB.Create(&relation).Error
	return
}

func DeleteRalation(idFrom, idTo int64) (err error) {
	err = dao.DB.Where("id_from=? and id_to=?", idFrom, idTo).Delete(&Comment{}).Error
	return
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	log, err := models.GetAUserInfoByEmail(aLoginUser)
	if err != nil {
		fmt.Println(err)
	}
	idTo := c.Query("to_user_id")
	if token == aLoginUser+log.Pwd {
		idT, _ := strconv.ParseInt(idTo, 10, 64)
		relation := Relation{
			IdFrom: log.Id,
			IdTo:   idT,
		}
		Type := c.Query("action_type")
		if Type == string(rune(1)) {
			err = CreateRalation(&relation)
		} else {
			err = DeleteRalation(log.Id, idT)
		}
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
