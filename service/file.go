// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 20:47
package service

import (
	"mime/multipart"
	"os"

	"cloud-disk/config"
	"cloud-disk/log"
	"cloud-disk/model"
	"cloud-disk/model/disk"
)

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

func UploadFileToModel(files []*multipart.FileHeader) error {
	for _, file := range files {
		err := model.AddFile(model.File{
			Name:        file.Filename,
			Uid:         config.NowUser.Uid,
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

func UploadFileToDisk(files []*multipart.FileHeader) error {
	return disk.UploadFile(files)
}

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

func DeleteFileToModel(name string, path string, isdir bool) error {
	return model.DelFile(model.File{
		Name:        name,
		Path:        path,
		IsDirectory: false,
	})
}

func DeleteFileToDisk(name string, path string, isdir bool) error {
	return disk.DeleteFile(name, path, isdir)
}

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
