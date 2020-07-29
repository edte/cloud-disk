// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 12:39
package config

type logFileConfig struct {
	Path string
	Name string
}

var LogFileConfig = logFileConfig{
	Path: "log/",
	Name: "log",
}

type databaseConfig struct {
	Name     string
	Type     string
	User     string
	Password string
	Host     string
	Port     int
}

var DatabaseConfig = databaseConfig{
	Name:     "cloud-disk",
	Type:     "mysql",
	User:     "root",
	Password: "mima",
	Host:     "127.0.0.1",
	Port:     3306,
}
