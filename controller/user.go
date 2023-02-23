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

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

var (
	aLoginUser = "none"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//var userInfo models.UserInfo
	token := username + password
	//userInfo := models.UserInfo{username, password}
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	_, exist := usersLoginInfo[token]
	fmt.Println(exist)
	getUser, getErr := GetAUserByName(username)
	if getErr != nil {
		fmt.Println(getErr)
	}
	fmt.Printf("getUser:%#v\n", getUser)
	if getUser != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		listUser, countErr := GetAllUser()
		if countErr != nil {
			fmt.Println(countErr)
		}
		fmt.Printf("listUser:%#v\n", listUser)
		count := len(listUser)

		fmt.Println(count)
		fmt.Println("count: " + string(rune(count)))
		var userIdSequence = int64(count)
		atomic.AddInt64(&userIdSequence, 1)
		newStoreUser := User{
			Id:            userIdSequence,
			Name:          username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		newUserInfo := models.UserInfo{
			Id:    userIdSequence,
			Email: username,
			Pwd:   password,
		}
		e := CreateAUser(&newStoreUser)
		e = models.CreateAUserInfo(&newUserInfo)
		if e != nil {
			panic(e)
		}
		usersLoginInfo[token] = newStoreUser
		aLoginUser = username
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

	//var loginUser models.UserInfo
	token := username + password
	//err := c.BindJSON(&loginUser)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	user, exist := usersLoginInfo[token]
	fmt.Println(exist)
	getUser, getErr := GetAUserByName(username)
	getUserInfo, getErr := models.GetAUserInfoByEmail(username)
	if getUser == nil {
		fmt.Println(getErr)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"}})
		return
	}
	usersLoginInfo = map[string]User{
		getUserInfo.Email + getUserInfo.Pwd: *getUser,
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})

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
