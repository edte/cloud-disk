// @program: cloud-disk
// @author: edte
// @create: 2020-07-30 00:53
package service

import (
	"math/rand"
	"time"

	"github.com/skip2/go-qrcode"

	"cloud-disk/config"
	"cloud-disk/log"
	"cloud-disk/model"
)

// ShareForm 用于分享表单
type ShareForm struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	HasCode  bool              `json:"has_code"`
	Code     int64             `json:"code"`
	IsQrCode bool              `json:"is_qr_code"`
	Files    map[string]string `json:"files"`
}

// Share 分享文件
func Share(share ShareForm) (string, error) {
	url := GetUrl()
	var fileNames []string

	for _, s := range share.Files {
		fileNames = append(fileNames, s)
	}

	err := model.AddShare(model.Share{
		Name:        share.Name,
		Type:        share.Type,
		Uid:         NowUser.Uid,
		Count:       0,
		HasCode:     share.HasCode,
		Code:        share.Code,
		IsExpired:   false,
		ExpiredTime: GetTime(),
		IsQrcode:    share.IsQrCode,
		Url:         url,
	}, fileNames)

	return url, err
}

// GetTime 获取当前 1000 h 后的时间
func GetTime() time.Time {
	currentTime := time.Now()

	m, err := time.ParseDuration(config.DefaultExpiredTime)
	if err != nil {
		log.Begin().Errorf("failed get time:%s", err)
	}
	result := currentTime.Add(m)
	return result
}

func GetUrl() string {
	s := RandString()

	return config.RouterHost + "/file/share/" + s
}

// RandString get a rand string by 20 len
func RandString() string {
	n := 20

	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// GetQcCode
func GetQcCode(url string) ([]byte, error) {
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	png, err := q.PNG(256)
	if err != nil {
		return nil, err
	}

	return png, nil
}

// IsShareExist
func IsShareExist(url string) bool {
	return model.IsShareExist(config.RouterHost + "/file/share/" + url)
}

// GetShareFiles
func GetShareFiles(name string) []string {
	files, err := model.GetShareFiles(name)
	if err != nil {
		log.Begin().Errorf("failed get file by name:%s", err)
		return nil
	}

	var s []string
	for _, file := range files {
		s = append(s, file.FileName)
	}
	return s
}

// GetShareUrls 通过分享文件生成对应 url
func GetShareUrls(files []string) []string {
	var s []string

	for _, file := range files {
		s = append(s, config.RouterHost+"/file/download/"+file)
	}
	return s
}

// GetShareNameByUrl 获取分享名字
func GetShareNameByUrl(url string) (string, error) {
	return model.GetShareNameByUrl(config.RouterHost + "/file/share/" + url)
}

func GetExpiredTime(url string) (time.Time, error) {
	return model.GetExpiredTime(config.RouterHost + "/file/share/" + url)
}

// IsExpired 获取分享是否失效
func IsExpired(url string) (bool, error) {
	return model.IsExpired(config.RouterHost + "/file/share/" + url)
}
