// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 11:12
package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Uid      int
	Role     string
}

func AddUser(u *User) error {
	return DB.Create(&u).Error
}

func DelUser(u *User) error {
	return nil
}

func ListUsers() []User {
	return nil
}

func GetPasswd() {

}
