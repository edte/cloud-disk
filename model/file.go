// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 19:01
package model

import (
	"github.com/jinzhu/gorm"
)

// File 表示 file
type File struct {
	gorm.Model
	Name        string // file name
	Uid         string // owner
	Size        int64  // file size
	Path        string // where the file is int the disk
	Count       int    // download number
	Privacy     bool   // whether the file is privacy
	IsDirectory bool   // whether the file is directory
}

// CountDecrement 用于下载文件，下载数自增
func CountDecrement(fname string) error {
	f, err := GetFileByName(fname)
	if err != nil {
		return err
	}

	return DB.Model(&File{}).Where("name = ?", fname).Update("count", f.Count).Error
}

// AddFile 增加文件
func AddFile(f File) error {
	return DB.Create(&f).Error
}

// DelFile 删除文件
func DelFile(f File) error {
	return DB.Where("name = ? and path = ? and is_directory = ?", f.Name, f.Path, f.IsDirectory).Delete(&File{}).Error
}

// GetFileByName 获取文件信息
func GetFileByName(fname string) (File, error) {
	var f File
	err := DB.Where("name = ?", fname).First(&f).Error
	return f, err
}

func ModifyFile(f File) {

}
