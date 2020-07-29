// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 12:39
package config

import (
	"log"
	"os"
)

type logFileConfig struct {
	Path string
	Name string
}

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

var DatabaseConfig = databaseConfig{
	Name:     "cloud_disk",
	Type:     "mysql",
	User:     "root",
	Password: "mima",
	Host:     "127.0.0.1",
	Port:     3306,
}

type LoginForm struct {
	Username string
	Password string
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

type nowUser struct {
	Username string
	Password string
	Uid      string
}

// NowUser
var NowUser = nowUser{
}

// DefaultDiskPath
var DefaultDiskPath = "tmp/"
