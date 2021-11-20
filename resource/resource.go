package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// 全局配置文件
var HrmsConf *Config

// 分公司数据库映射表
var DbMapper = make(map[string]*gorm.DB)

// 默认DB，不作为业务使用
var DefaultDb *gorm.DB

type Gin struct {
	Port int64 `json:"port"`
}

// 解析cookie中的分公司Id，获取对应数据库实例
func HrmsDB(c *gin.Context) *gorm.DB {
	cookie, err := c.Cookie("user_cookie")
	if err != nil || cookie == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		return nil
	}
	branchId := strings.Split(cookie, "_")[2]
	dbName := fmt.Sprintf("hrms_%v", branchId)
	if db, ok := DbMapper[dbName]; ok {
		return db
	}
	c.HTML(http.StatusOK, "login.html", nil)
	return nil
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	DbName   string `json:"dbNname"`
}

type Mongo struct {
	IP      string `json:"ip"`
	Port    int64  `json:"port"`
	Dataset string `json:"dataset"`
}

var MongoClient *qmgo.Client

type Config struct {
	Gin   `json:"gin"`
	Db    `json:"db"`
	Mongo `json:"mongo"`
}
