// @program: cloud-disk
// @author: edte
// @create: 2020-07-30 00:14
package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"cloud-disk/log"
	"cloud-disk/router/response"
	"cloud-disk/service"
)

// ShareForm
type ShareForm struct {
	Key string `uri:"key"`
}

// Share
func Share(c *gin.Context) {
	var sf service.ShareForm

	if err := c.ShouldBindJSON(&sf); err != nil {
		log.Begin().Infof("share file form failed:%s", err)
		response.FormError(c)
		return
	}

	url, err := service.Share(sf)
	if err != nil {
		log.Begin().Errorf("share file failed:%s", err)
		response.Error(c, response.CodeTmp, "share file failed")
		return
	}

	if sf.IsQrCode {
		code, err := service.GetQcCode(url)
		if err != nil {
			log.Begin().Errorf("failed to get qccode:%s", err)
			response.Error(c, response.CodeTmp, "share file failed")
		}

		c.Writer.Header().Add("Content-Type", "image/png")
		c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", len(code)))
		c.Writer.Write(code)
	}
	response.OkWithData(c, url)
}

// HandleShare
func HandleShare(c *gin.Context) {
	var sf ShareForm

	if err := c.ShouldBindUri(&sf); err != nil {
		log.Begin().Infof("handle share file form failed:%s", err)
		response.FormError(c)
		return
	}

	// judge whether share is exist
	if !service.IsShareExist(sf.Key) {
		log.Begin().Infof("share not exist:%s")
		response.Error(c, response.CodeTmp, "share not exist")
		return
	}

	expired, err2 := service.IsExpired(sf.Key)
	if err2 != nil {
		log.Begin().Errorf("get isexpired failed:%s", err2)
	}

	if expired {
		response.OkWithData(c, "share has expired")
	}

	// get share name by share url
	name, err := service.GetShareNameByUrl(sf.Key)
	if err != nil {
		log.Begin().Errorf("get share name by url failed:%s", err)
	}
	// get share files
	files := service.GetShareFiles(name)
	// get share file urls
	urls := service.GetShareUrls(files)

	response.OkWithData(c, urls)
}
