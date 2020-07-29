// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 21:07
package disk

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"cloud-disk/config"
)

// UploadFile 上传文件
func UploadFile(files []*multipart.FileHeader) error {
	for _, file := range files {
		src, err := file.Open()
		if err != nil {

			return err
		}

		out, err := os.Create(config.DefaultDiskPath + file.Filename)
		defer out.Close()
		if err != nil {

			return err
		}

		if _, err := io.Copy(out, src); err != nil {
			fmt.Println("aabbcc")
			return err
		}
	}

	return nil
}

// todo: os.remove 同时可以删除文件和目录，而 os.create 和 os.mkdir 似乎不能创建同名文件和文件夹，所以这里有待解决
// DeleteFile 删除文件
func DeleteFile(name string, path string, isdir bool) error {
	return os.Remove(path + "/" + name)
}

// DownloadFile 获取文件
func DownloadFile(path string, name string) (*os.File, error) {
	return os.Open(path + "/" + name)
}
