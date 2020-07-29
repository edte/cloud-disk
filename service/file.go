// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 20:47
package service

import (
	"mime/multipart"
	"os"

	"cloud-disk/log"
	"cloud-disk/model"
	"cloud-disk/model/disk"
)

// UploadFile 用于上传文件
func UploadFile(files []*multipart.FileHeader) error {
	if err := UploadFileToModel(files); err != nil {
		log.Begin().Errorf("failed to upload file to database:%s", err)
		return err
	}

	if err := UploadFileToDisk(files); err != nil {
		log.Begin().Errorf("failed to upload file to disk:%s", err)
		return err
	}
	return nil
}

// UploadFileToModel 用于上传文件到 database
func UploadFileToModel(files []*multipart.FileHeader) error {
	for _, file := range files {
		err := model.AddFile(model.File{
			Name:        file.Filename,
			Uid:         NowUser.Uid,
			Size:        file.Size,
			Path:        "tmp",
			Count:       0,
			Privacy:     false,
			IsDirectory: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// UploadFileToDisk 上传文件到 disk
func UploadFileToDisk(files []*multipart.FileHeader) error {
	return disk.UploadFile(files)
}

// DeleteFile 删除文件
func DeleteFile(name string, path string, isdir bool) error {
	if err := DeleteFileToModel(name, path, isdir); err != nil {
		log.Begin().Errorf("failed to delete file at model:%s", err)
		return err
	}
	if err := DeleteFileToDisk(name, path, isdir); err != nil {
		log.Begin().Errorf("failed to delete file at disk:%s", err)
		return err
	}
	return nil
}

// DeleteFileToModel 从 database 删除文件
func DeleteFileToModel(name string, path string, isdir bool) error {
	return model.DelFile(model.File{
		Name:        name,
		Path:        path,
		IsDirectory: false,
	})
}

// DeleteFileToDisk 从 disk 删除文件
func DeleteFileToDisk(name string, path string, isdir bool) error {
	return disk.DeleteFile(name, path, isdir)
}

// DownloadFile 下载文件
func DownloadFile(path string, name string) (*os.File, error) {
	if err := model.CountDecrement(name); err != nil {
		return nil, err
	}
	file, err := disk.DownloadFile(path, name)
	if err != nil {
		log.Begin().Errorf("failed download file:%s", err)
		return nil, err
	}
	return file, nil
}
