package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/models"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user models.User
	token := username + password
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, exist := usersLoginInfo[token]
	fmt.Println(exist)
	getUser, getErr := models.GetAUser(username)
	if getErr != nil {
		panic(getErr)
	}
	if getUser != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		listUser, countErr := models.GetAllUser()
		if countErr != nil {
			panic(countErr)
		}
		count := len(listUser)
		atomic.AddInt64(&userIdSequence, int64(count)+1)
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		e := models.CreateAUser(&user)
		if e != nil {
			panic(e)
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var loginUser models.User
	token := username + password
	err := c.BindJSON(&loginUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, exist := usersLoginInfo[token]
	fmt.Println(exist)
	getUser, getErr := models.GetAUser(username)
	if getErr != nil {
		fmt.Println(getErr)
		return
	}
	if getUser != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
