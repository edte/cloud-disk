// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 14:18
package handlers

import (
	"github.com/gin-gonic/gin"

	"cloud-disk/service"
)

func Login(c *gin.Context) {
	service.Login()
}

func Register(c *gin.Context) {
	service.Register()
}
