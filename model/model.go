// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 11:54
package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"cloud-disk/config"
	"cloud-disk/log"
)

var DB *gorm.DB

func InitModel() {
	log.Begin().Info("begin init database..")

	db, err := gorm.Open(config.DatabaseConfig.Type,
		fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DatabaseConfig.User,
			config.DatabaseConfig.Password,
			config.DatabaseConfig.Host,
			config.DatabaseConfig.Port,
			config.DatabaseConfig.Name))

	if err != nil {
		log.Begin().Fatalf("failed to connect database:%v", err)
	}
	DB = db
	setTable()
	addDefaultData()
	log.Begin().Info("init database successful")
}

func setTable() {
	log.Begin().Info("begin init table...")
	// user
	if DB.HasTable(&User{}) {
		DB.AutoMigrate(&User{})
	} else {
		DB.CreateTable(&User{})
	}
	// file
	if DB.HasTable(&File{}) {
		DB.AutoMigrate(&File{})
	} else {
		DB.CreateTable(&File{})
	}

	log.Begin().Info("init table successful")
}

func addDefaultData() {
	log.Begin().Info("begin set default table data...")
	u := User{
		Username: "root",
		Password: "root",
		Uid:      "0",
		Role:     "root",
	}
	log.Begin().Info("begin add root user..")
	if AddUser(&u) != nil {
		log.Begin().Errorf("failed to add root user")
	}
	log.Begin().Info("add root user successful")

	log.Begin().Info("set default table data successful")
}
