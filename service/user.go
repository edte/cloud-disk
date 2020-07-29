// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 13:52
package service

import (
	"github.com/gofrs/uuid"

	"cloud-disk/log"
	"cloud-disk/model"
)

type nowUser struct {
	Username string
	Password string
	Uid      string
}

// NowUser 用于备份当用户信息
var NowUser = nowUser{
}

// LoginForm 用于登录表单
type LoginForm struct {
	Username string
	Password string
}

// IsUserExist 用于判断用户是否存在
func IsUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Begin().Info("user not exist:%v", err)
		return false
	}
	return true
}

// IsPasswdOk 用于判断密码是否正确
func IsPasswdOk(l LoginForm) bool {
	u, err := model.GetUserByUsername(l.Username)
	if err != nil {
		log.Begin().Error("get user by username failed:%v", err)
	}
	return u.Password == l.Password
}

// AddUser 用于增加用户
func AddUser(l LoginForm, uid string) error {
	err := model.AddUser(&model.User{
		Username: l.Username,
		Password: l.Password,
		Uid:      uid,
		Role:     "general",
	})
	return err
}

// IsRegister 用于判断是否注册
func IsRegister(l LoginForm) bool {
	u, err := model.GetUserByUsername(l.Username)
	if err != nil {
		log.Begin().Infof("Determine whether the registration:%s", err)
		return false
	}
	return u.Password == l.Password
}

// GetARandomUid 获取随机的 uid
func GetARandomUid() string {
	v4, err := uuid.NewV4()
	if err != nil {
		log.Begin().Fatalf("failed to get a random uuid:%v", err)
	}
	return v4.String()
}

// GetUid
func GetUid(username string) string {
	u, err := model.GetUserByUsername(username)
	if err != nil {
		log.Begin().Errorf("failed to get uid:%v", err)
	}
	return u.Uid
}
