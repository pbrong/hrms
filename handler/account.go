package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"log"
	"net/http"
	"strings"
)

func Ping(c *gin.Context) {
	c.HTML(http.StatusOK, "staff_manage.html", gin.H{
		"create": true,
	})
}

func Index(c *gin.Context) {
	// 判断是否存在cookie
	cookie, err := c.Cookie("user_cookie")
	if err != nil || cookie == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	// 已登陆
	user := strings.Split(cookie, "_")
	userType := user[0]
	userNo := user[1]
	userName := user[3]
	c.HTML(http.StatusOK, "index.html", gin.H{
		//"title":     fmt.Sprintf("欢迎%v:%v登陆HRMS", userType, userNo),
		"title":      fmt.Sprintf("人力资源管理系统"),
		"user_type":  userType,
		"staff_id":   userNo,
		"staff_name": base64Decode(userName),
	})
}

func base64Decode(name string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(name)
	if err != nil {
		log.Fatalln(err)
		return "企业员工"
	}
	return string(decodeBytes)
}

func RenderAuthority(c *gin.Context) {
	cookie, err := c.Cookie("user_cookie")
	if err != nil || cookie == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	modelName := c.Param("modelName")
	userType := strings.Split(cookie, "_")[0]
	dto := &model.GetAuthorityDetailDTO{
		UserType: userType,
		Model:    modelName,
	}
	autoContent, err := service.GetAuthorityDetailByUserTypeAndModel(c, dto)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	autoMap := make(map[string]bool)
	autoList := strings.Split(autoContent, "|")
	for _, autority := range autoList {
		autoMap[autority] = true
	}
	//c.JSON(200, autoMap)
	c.HTML(http.StatusOK, modelName+".html", autoMap)
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
	dbName := fmt.Sprintf("hrms_%v", loginR.BranchId)
	log.Printf("[login db name = %v]", dbName)
	var hrmsDB *gorm.DB
	var ok bool
	if hrmsDB, ok = resource.DbMapper[dbName]; !ok {
		log.Printf("[Login err, 无法获取到该分公司db名称, name = %v]", dbName)
		c.JSON(200, gin.H{
			"status": 5000,
			"result": fmt.Sprintf("[Login err, 无法获取到该分公司db名称, name = %v]", dbName),
		})
		return
	}
	log.Printf("[handler.Login] login R = %v", loginR)
	var loginDb model.Authority
	var staff model.Staff
	hrmsDB.Where("staff_id = ? and user_password = ?",
		loginR.UserNo, service.MD5(loginR.UserPassword)).First(&loginDb)
	if loginDb.StaffId != loginR.UserNo {
		log.Printf("[handler.Login] user login fail, user = %v", loginR)
		c.JSON(200, gin.H{
			"status": 2001,
			"result": "check fail",
		})
		return
	}
	hrmsDB.Where("staff_id = ?", loginDb.StaffId).Find(&staff)

	log.Printf("[handler.Login] user login success, user = %v", loginR)
	// set cookie user_cookie=角色_工号_分公司ID_员工姓名(base64编码)
	c.SetCookie("user_cookie", fmt.Sprintf("%v_%v_%v_%v", loginDb.UserType, loginDb.StaffId, loginR.BranchId,
		base64.StdEncoding.EncodeToString([]byte(staff.StaffName))), 0, "/", "*", false, false)

	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func Quit(c *gin.Context) {
	c.SetCookie("user_cookie", "null", -1, "/", "*", false, false)
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

//func Quit(c *gin.Context) {
//	var quitR model.LoginDTO
//	if err := c.ShouldBindJSON(&quitR); err != nil {
//		log.Printf("[handler.Quit] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5001,
//			"result": err.Error(),
//		})
//		return
//	}
//	var quitDb model.Authority
//	resource.HrmsDB(c).Where("staff_id = ?",
//		quitR.UserNo).First(&quitDb)
//	if quitDb.UserType == "" || quitDb.StaffId == "" {
//		log.Printf("[handler.Quit] user quit fail, user = %v", quitR)
//		c.JSON(200, gin.H{
//			"status": 5000,
//		})
//	}
//	log.Printf("[handler.Quit] user quit success, user = %v", quitR)
//	// del cookie user_cookie
//	c.SetCookie("user_cookie", fmt.Sprintf("%v_%v", quitDb.UserType, quitDb.StaffId), -1, "/", "*", false, false)
//	c.JSON(200, gin.H{
//		"status": 2000,
//	})
//}
