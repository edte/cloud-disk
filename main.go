// @program: cloud-disk
// @author: edte
// @create: 2020-07-28 21:33
package main

import (
	"cloud-disk/model"
	"cloud-disk/router"
	"cloud-disk/service"
)

func main() {
	model.InitModel()
	service.InitService()
	router.InitRouter()
}
