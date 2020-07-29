// @program: cloud-disk
// @author: edte
// @create: 2020-07-30 00:13
package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Share 表示一次分享
type Share struct {
	gorm.Model
	Name        string    // share name
	Type        string    // share type: file,dir,all
	Uid         string    // share uid
	Count       int64     // download number
	HasCode     bool      // whether has password
	Code        int64     // share password
	IsExpired   bool      // whether is expired
	ExpiredTime time.Time // when it expired
	IsQrcode    bool      // Is qrcode
	Url         string    // share url
}

// ShareFiles 表示一次分享的文件
type ShareFiles struct {
	gorm.Model
	Name     string // share name
	FileName string // share file name include file path
}

// AddShare 增加 share
func AddShare(share Share, files []string) error {
	if DB.Create(&share).Error != nil {
		return DB.Create(&share).Error
	}

	for _, file := range files {
		var sf ShareFiles
		sf.Name = share.Name
		sf.FileName = file

		if err := DB.Create(&sf).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetShareFiles 获取分享的文件
func GetShareFiles(shareName string) (files []ShareFiles, err error) {
	errs := DB.Where("name = ?", shareName).Find(&files).GetErrors()
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

// IsShareExist 判断分享是否存在
func IsShareExist(url string) bool {
	fmt.Println(url)
	var s Share
	err := DB.Where("url = ?", url).First(&s).Error
	return err == nil
}

// GetShareNameByUrl
func GetShareNameByUrl(url string) (string, error) {
	var s Share
	err := DB.Where("url = ?", url).First(&s).Error
	if err != nil {
		return "", err
	}
	return s.Name, nil
}

// GetExpiredTime
func GetExpiredTime(url string) (time.Time, error) {
	var s Share
	err := DB.Where("url = ?", url).First(&s).Error
	if err != nil {
		return time.Time{}, err
	}
	return s.ExpiredTime, nil
}

// IsExpired 用于获取分享是否失效
func IsExpired(url string) (bool, error) {
	var s Share
	err := DB.Where("url = ?", url).First(&s).Error
	if err != nil {
		return false, err
	}
	return s.IsExpired, nil
}
