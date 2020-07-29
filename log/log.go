// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 12:53
package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"cloud-disk/config"
)

func Begin() *logrus.Logger {
	logFilePath := config.LogFileConfig.Path
	logFileName := config.LogFileConfig.Name

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Printf("failed to open log file: %v\n", err)
	}
	fmt.Println(err)
	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	// 设置方法名
	logger.SetReportCaller(true)

	// 同时打印日志和存到文件中
	logger.SetOutput(io.MultiWriter(os.Stdout, src))

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
