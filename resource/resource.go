package resource

import "gorm.io/gorm"

// 全局配置文件
var HrmsConf *Config
// 全局GormDB
var HrmsDB *gorm.DB

type Gin struct {
	Port int64 `json:"port"`
}

type Db struct {
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port int64 	`json:"port"`
	DbName string `json:"dbNname"`
}

type Config struct {
	Gin `json:"gin"`
	Db `json:"db"`
}
