// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 13:52
package service

import (
	"github.com/gofrs/uuid"

	"cloud-disk/config"
	"cloud-disk/log"
	"cloud-disk/model"
)

func IsUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Begin().Info("user not exist:%v", err)
		return false
	}
	return true
}

func IsPasswdOk(l config.LoginForm) bool {
	u, err := model.GetUserByUsername(l.Username)
	if err != nil {
		log.Begin().Error("get user by username failed:%v", err)
	}
	return u.Password == l.Password
}

func AddUser(l config.LoginForm, uid string) error {
	err := model.AddUser(&model.User{
		Username: l.Username,
		Password: l.Password,
		Uid:      uid,
		Role:     "general",
	})
	return err
}

func IsRegister(l config.LoginForm) bool {
	u, err := model.GetUserByUsername(l.Username)
	if err != nil {
		log.Begin().Infof("Determine whether the registration:%s", err)
		return false
	}
	return u.Password == l.Password
}

func GetARandomUid() string {
	v4, err := uuid.NewV4()
	if err != nil {
		log.Begin().Fatalf("failed to get a random uuid:%v", err)
	}
	return v4.String()
}

func GetUid(username string) string {
	u, err := model.GetUserByUsername(username)
	if err != nil {
		log.Begin().Errorf("failed to get uid:%v", err)
	}
	return u.Uid
}
