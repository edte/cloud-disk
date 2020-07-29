// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 11:12
package model

import (
	"github.com/jinzhu/gorm"
)

// User 表示用户
type User struct {
	gorm.Model
	Username string // username
	Password string // password
	Uid      string // uid
	Role     string // user role
}

// AddUser 增加用户
func AddUser(u *User) error {
	return DB.Create(&u).Error
}

// DelUser 删除用户
func DelUser(u *User) error {
	return DB.Delete(&u).Error
}

// ListUsers 列出所有用户
func ListUsers() []User {
	return nil
}

// GetUserByUsername 获取用户信息
func GetUserByUsername(username string) (User, error) {
	var u User

	err := DB.Where("username = ?", username).First(&u).Error
	return u, err
}
