// @program: cloud-disk
// @author: edte
// @create: 2020-07-29 10:51
package router

import (
	"github.com/gin-gonic/gin"

	"cloud-disk/log"
	"cloud-disk/router/handlers"
	"cloud-disk/router/middleware"
)

func InitRouter() {
	log.Begin().Info("begin init router...")

	r := gin.Default()

	setupRouter(r)

	err := r.Run()
	if err != nil {
		log.Begin().Fatalf("failed to init router")
	}

	log.Begin().Info("init router successful")
}

func setupRouter(r *gin.Engine) {
	r.Use(middleware.LoggerToFile())

	user := r.Group("/user")
	{
		user.POST("/login", handlers.Login)
		user.POST("/register", handlers.Register)
	}

	file := r.Group("/file")
	{
		file.POST("/upload")
		file.DELETE("/delete")
		file.GET("/download/:fid")
		file.GET("/share/:fid")
	}
}
