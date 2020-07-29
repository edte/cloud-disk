// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 10:51
package router

import (
	"github.com/gin-gonic/gin"

	"cloud-disk/config"
	"cloud-disk/log"
	"cloud-disk/router/handlers"
	"cloud-disk/router/middleware"
)

func InitRouter() {
	log.Begin().Info("begin init router...")

	r := gin.Default()

	setupRouter(r)

	err := r.Run(config.RouterHost)
	if err != nil {
		log.Begin().Fatalf("failed to init router")
	}

	log.Begin().Info("init router successful")
}

func setupRouter(r *gin.Engine) {
	// log record，解决跨域
	r.Use(middleware.LoggerToFile(), middleware.Cors())

	// user
	user := r.Group("/user")
	{
		user.POST("/login", handlers.Login)
		user.POST("/register", handlers.Register)
	}

	// file
	file := r.Group("/file", middleware.AuthRequired())
	{
		file.POST("/upload", handlers.Upload)

		publicFileRequired := file.Group("", middleware.PublicFile())
		{
			publicFileRequired.DELETE("/delete", handlers.Delete)
			publicFileRequired.GET("/download/:path/:name", handlers.Download)
			publicFileRequired.GET("/share", handlers.Share)
			publicFileRequired.GET("/share/:key", handlers.HandleShare)
		}
	}
}
