// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 12:39
package config

import (
	"log"
	"os"
)

// DefaultDiskPath 用来表示 file 的默认路径
var DefaultDiskPath = "tmp/"

type logFileConfig struct {
	Path string
	Name string
}

// LogFileConfig 用于存日志信息
var LogFileConfig = logFileConfig{
	Path: "",
	Name: "log.txt",
}

type databaseConfig struct {
	Name     string
	Type     string
	User     string
	Password string
	Host     string
	Port     int
}

// DatabaseConfig 用于数据库连接
var DatabaseConfig = databaseConfig{
	Name:     "cloud_disk",
	Type:     "mysql",
	User:     "root",
	Password: "mima",
	Host:     "127.0.0.1",
	Port:     3306,
}

type cookieConfig struct {
	Name     string
	Value    string
	Host     string
	Path     string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

// CookieConfig 用于 cookie 配置
var CookieConfig = cookieConfig{
	Name:     "soidfjosd",
	Value:    "",
	Host:     "localhost:8080",
	Path:     "/",
	MaxAge:   10000,
	Secure:   false,
	HttpOnly: true,
}

// JwtSecret
var JwtSecret = "osdjfowjeogjweoi"

// InitConfig
func InitConfig() {
	_, err := os.Create(LogFileConfig.Path + LogFileConfig.Name)
	if err != nil {
		log.Fatalf("failed init log file")
	}

	err = os.Mkdir("tmp/", 0777)
	if os.IsExist(err) {
		os.Remove("tmp")
		os.Mkdir("tmp/", 0777)
	}
	log.Println(err)
}

// RouterHost 表示 router 的 host
var RouterHost = "127.0.0.1:8080"

// DefaultExpiredTime 表示 share 默认失效时间
var DefaultExpiredTime = "1000h"
