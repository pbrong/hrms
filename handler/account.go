package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
	"net/http"
	"strings"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Index(c *gin.Context) {
	// 判断是否存在cookie
	cookie, err := c.Cookie("user_cookie")
	if err != nil || cookie == "" {
		c.HTML(http.StatusOK, "login-1.html", nil)
		return
	}
	// 已登陆
	user := strings.Split(cookie, "_")
	userType := user[0]
	userNo := user[1]
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     fmt.Sprintf("欢迎%v:%v登陆HRMS", userType, userNo),
		"user_type": userType,
		"staff_id":  userNo,
	})
}

func Login(c *gin.Context) {
	var loginR model.LoginDTO
	if err := c.ShouldBindJSON(&loginR); err != nil {
		log.Printf("[handler.Login] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	log.Printf("[handler.Login] login R = %v", loginR)
	var loginDb model.Login
	resource.HrmsDB.Where("staff_id = ? and user_password = ?",
		loginR.UserNo, loginR.UserPassword).First(&loginDb)
	if loginDb.StaffId != loginR.UserNo {
		log.Printf("[handler.Login] user login fail, user = %v", loginR)
		c.JSON(200, gin.H{
			"status": 2001,
			"result": "check fail",
		})
		return
	}
	log.Printf("[handler.Login] user login success, user = %v", loginR)
	// set cookie user_cookie=sys_3117000001
	c.SetCookie("user_cookie", fmt.Sprintf("%v_%v", loginDb.UserType, loginDb.StaffId), 0, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func Quit(c *gin.Context) {
	var quitR model.LoginDTO
	if err := c.ShouldBindJSON(&quitR); err != nil {
		log.Printf("[handler.Quit] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	var quitDb model.Login
	resource.HrmsDB.Where("staff_id = ?",
		quitR.UserNo).First(&quitDb)
	if quitDb.UserType == "" || quitDb.StaffId == "" {
		log.Printf("[handler.Quit] user quit fail, user = %v", quitR)
		c.JSON(200, gin.H{
			"status": 5000,
		})
	}
	log.Printf("[handler.Quit] user quit success, user = %v", quitR)
	// del cookie user_cookie
	c.SetCookie("user_cookie", fmt.Sprintf("%v_%v", quitDb.UserType, quitDb.StaffId), -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status": 2000,
	})
}
