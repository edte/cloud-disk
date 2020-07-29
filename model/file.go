// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 19:01
package model

import (
	"github.com/jinzhu/gorm"
)

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

func CountDecrement(fname string) error {
	f, err := GetFileByName(fname)
	if err != nil {
		return err
	}

	return DB.Model(&File{}).Where("name = ?", fname).Update("count", f.Count).Error
}

func AddFile(f File) error {
	return DB.Create(&f).Error
}

func DelFile(f File) error {
	return DB.Where("name = ? and path = ? and is_directory = ?", f.Name, f.Path, f.IsDirectory).Delete(&File{}).Error
}

func GetFileByName(fname string) (File, error) {
	var f File
	err := DB.Where("name = ?", fname).First(&f).Error
	return f, err
}

func ModifyFile(f File) {

}
