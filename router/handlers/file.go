// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 19:25
package handlers

import (
	"io"

	"github.com/gin-gonic/gin"

	"cloud-disk/log"
	"cloud-disk/router/response"
	"cloud-disk/service"
)

// FileForm 用于登录表单
type FileForm struct {
	Name  string
	Isdir bool
	Path  string
}

// DownloadForm 用于下载表单
type DownloadForm struct {
	Name string `uri:"name" binding:"required"`
	Path string `uri:"path" binding:"required"`
}

// Upload 用于上传文件
func Upload(c *gin.Context) {
	log.Begin().Infof("upload file begin...")

	// 设置最大支持上传的文件，这里表示 20 m，原因暂时看不懂
	if c.Request.ParseMultipartForm(20<<20) != nil {
		response.Error(c, response.CodeTmp, "too big files")
		return
	}

	formData := c.Request.MultipartForm
	files := formData.File["file"]

	if service.UploadFile(files) == nil {
		response.OkWithData(c, "upload successful")
		log.Begin().Infof("upload file successful")
		return
	}
	response.Error(c, response.CodeTmp, "upload file failed")
}

// Delete 用于删除文件
func Delete(c *gin.Context) {
	log.Begin().Infof("delete file begin...")
	var ff FileForm

	if err := c.ShouldBindJSON(&ff); err != nil {
		response.FormError(c)
		return
	}

	if service.DeleteFile(ff.Name, ff.Path, ff.Isdir) == nil {
		response.OkWithData(c, "delete file successful")
		return
	}

	response.OkWithData(c, "delete file failed")
}

// Download 用于下载文件
func Download(c *gin.Context) {
	log.Begin().Infof("download begin...")
	var df DownloadForm

	// 这里只有两层 uri，即 path/file，如果 path 有多层的话，发送请求的是否需要对 / 转码
	if err := c.ShouldBindUri(&df); err != nil {
		log.Begin().Infof("download file form failed:%s", err)
		response.FormError(c)
		return
	}

	f, err := service.DownloadFile(df.Path, df.Name)
	if err != nil {
		response.Error(c, response.CodeTmp, "download failed")
		return
	}

	// 设置这个 header 表示默认传输 byte stream，browser 默认下载
	c.Writer.Header().Add("Content-type", "application/octet-stream")

	_, err = io.Copy(c.Writer, f)
	if err != nil {
		response.Error(c, response.CodeTmp, "download failed")
		log.Begin().Errorf("failed to download:%s", err)
	}
	log.Begin().Infof("download successful...")
}
