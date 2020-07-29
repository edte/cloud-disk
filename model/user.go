// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 11:12
package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string // username
	Password string // password
	Uid      string // uid
	Role     string // user role
}

func AddUser(u *User) error {
	return DB.Create(&u).Error
}

func DelUser(u *User) error {
	return DB.Delete(&u).Error
}

func ListUsers() []User {
	return nil
}

func GetUserByUsername(username string) (User, error) {
	var u User

	err := DB.Where("username = ?", username).First(&u).Error
	return u, err
}
