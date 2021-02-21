package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hrms/handler"
	"hrms/resource"
	"log"
	"net/http"
	"os"
)

func InitConfig() error {
	config := &resource.Config{}
	vip := viper.New()
	vip.AddConfigPath("./config")
	vip.SetConfigType("yaml")
	// 环境判断
	env := os.Getenv("HRMS_ENV")
	if env == "" || env == "dev" {
		// 开发环境
		vip.SetConfigName("config-dev")
	}
	if env == "prod" {
		// 生产环境
		vip.SetConfigName("config-prod")
	}
	err := vip.ReadInConfig()
	if err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}
	if err := vip.Unmarshal(config); err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}
	log.Printf("[config.Init] 初始化配置成功,config=%v", config)
	resource.HrmsConf = config
	return nil
}

func InitGin() error {
	server := gin.Default()
	// 静态资源及模板配置
	htmlInit(server)
	// 初始化路由
	routerInit(server)
	err := server.Run(fmt.Sprintf(":%v", resource.HrmsConf.Gin.Port))
	if err != nil {
		log.Printf("[InitGin] err = %v", err)
	}
	log.Printf("[InitGin] success")
	return err
}

func routerInit(server *gin.Engine) {
	// 测试
	server.GET("/ping", handler.Ping)
	// 首页重定向
	server.GET("/index", handler.Index)
	// 账户相关
	accountGroup := server.Group("/account")
	accountGroup.POST("/login", handler.Login)
	accountGroup.POST("/quit", handler.Quit)
	// 部门相关
	departGroup := server.Group("/depart")
	departGroup.POST("/create", handler.DepartCreate)
	departGroup.DELETE("/del/:dep_id", handler.DepartDel)
	departGroup.POST("/edit", handler.DepartEdit)
	departGroup.GET("/query/:dep_id", handler.DepartQuery)
	// 职级相关
	rankGroup := server.Group("/rank")
	rankGroup.POST("/create", handler.RankCreate)
	rankGroup.DELETE("/del/:rank_id", handler.RankDel)
	rankGroup.POST("/edit", handler.RankEdit)
	rankGroup.GET("/query/:rank_id", handler.RankQuery)
}

func htmlInit(server *gin.Engine) {
	// 静态资源
	server.StaticFS("/static", http.Dir("./static"))
	server.StaticFS("/views", http.Dir("./views"))
	// HTML模板加载
	server.LoadHTMLGlob("views/*")
	// 404页面
	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
}

func InitGorm() error {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		resource.HrmsConf.Db.User,
		resource.HrmsConf.Db.Password,
		resource.HrmsConf.Db.Host,
		resource.HrmsConf.Db.Port,
		resource.HrmsConf.Db.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[InitGorm] err = %v", err)
		return err
	}
	resource.HrmsDB = db
	log.Printf("[InitGorm] success")
	return nil
}
func main() {
	if err := InitConfig(); err != nil {
		panic(err)
	}
	if err := InitGorm(); err != nil {
		panic(err)
	}
	if err := InitGin(); err != nil {
		panic(err)
	}
}
