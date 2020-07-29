// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 14:18
package handlers

import (
	"github.com/gin-gonic/gin"

	"cloud-disk/config"
	"cloud-disk/log"
	"cloud-disk/router/response"
	"cloud-disk/service"
	"cloud-disk/service/jwt"
)

// Login 用于登录
func Login(c *gin.Context) {
	log.Begin().Infof("begin login...")

	if hasToken(c) {
		if isTokenOk(c) {
			backUpUser(c)
			response.OkWithData(c, "you has login")
			return
		}
		log.Begin().Info("token not correct")
		response.Error(c, response.CodeTmp, "token not correct")
		return
	}

	var l service.LoginForm

	if err := c.ShouldBindJSON(&l); err != nil {
		response.FormError(c)
		return
	}

	if !service.IsUserExist(l.Username) {
		response.Error(c, response.CodeTmp, "user is not exist")
		return
	}

	if !service.IsPasswdOk(l) {
		log.Begin().Info("password not right")
		response.Error(c, response.CodeTmp, "password not right")
		return
	}

	token, err := jwt.GenerateToken(l.Username, l.Password, service.GetUid(l.Username))
	if err != nil {
		log.Begin().Errorf("failed to generate token:%s", err)
		return
	}

	config.CookieConfig.Value = token

	SetCookie(c)

	response.OkWithData(c, "login successful")

	log.Begin().Info("login successful")
}

func isTokenOk(c *gin.Context) bool {
	value, _ := c.Cookie(config.CookieConfig.Name)
	_, err := jwt.ParseToken(value)
	return err == nil
}

func hasToken(c *gin.Context) bool {
	_, err := c.Cookie(config.CookieConfig.Name)
	return err == nil
}

// HasLogin 用于判断是否注册
func HasLogin(c *gin.Context) bool {
	if !hasToken(c) {
		return false
	}
	if !isTokenOk(c) {
		log.Begin().Info("token not correct")
		response.Error(c, response.CodeTmp, "token not correct")
		return false
	}
	return true
}

func backUpUser(c *gin.Context) {
	value, _ := c.Cookie(config.CookieConfig.Name)

	user, _ := jwt.ParseToken(value)
	service.NowUser.Username = user.Username
	service.NowUser.Password = user.Password
	service.NowUser.Uid = user.Uid
}

// Register 注册
func Register(c *gin.Context) {
	var l service.LoginForm

	if err := c.ShouldBindJSON(&l); err != nil {
		response.FormError(c)
		return
	}

	if service.IsRegister(l) {
		response.Error(c, response.CodeTmp, "user is exist")
		return
	}

	uid := service.GetARandomUid()
	if err := service.AddUser(l, uid); err != nil {
		log.Begin().Errorf("failed to add user when register:%v", err)
		return
	}

	token, err := jwt.GenerateToken(l.Username, l.Password, uid)
	if err != nil {
		log.Begin().Errorf("failed to generate token:%s", err)
		return
	}

	config.CookieConfig.Value = token

	// todo: 这里直接把 token 放 cookie 里了，避免测试时需要手动设置 auth 和设置获取 token 的 middleware
	SetCookie(c)

	service.NowUser.Username = l.Username
	service.NowUser.Password = l.Password
	service.NowUser.Uid = uid

	response.OkWithData(c, "register successful!")
}

// SetCookie 设置 cookie
func SetCookie(c *gin.Context) {
	c.SetCookie(
		config.CookieConfig.Name,
		config.CookieConfig.Value,
		config.CookieConfig.MaxAge,
		config.CookieConfig.Path,
		config.CookieConfig.Host,
		config.CookieConfig.Secure,
		config.CookieConfig.HttpOnly)
}
