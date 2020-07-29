// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 14:21
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 业务码
const (
	// 正常响应
	CodeOk = 10000
	// 请求表单错误
	CodeFormError = 10001
	// todo: 设计业务码
	CodeTmp = 10002
)

// Ok
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": CodeOk})
}

// FormError
func FormError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": CodeFormError, "message": "request form error!"})
}

// OkWithData
func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": CodeOk, "message": "ok", "data": data})
}

// Error
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
}
