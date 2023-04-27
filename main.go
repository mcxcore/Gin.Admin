package main

import (
	"Gin.Admin/controllers"
	"Gin.Admin/models"
	"Gin.Admin/pkg/gredis"
	"Gin.Admin/pkg/logging"
	"Gin.Admin/pkg/setting"
	"Gin.Admin/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	引入包

gorm    操作数据库
gredis   操作缓存
......
logger  日志包
jwt   包
swagger 包
*/
func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/go", controllers.Lets)
	server := &http.Server{
		Addr:    ":8001",
		Handler: r,
	}
	server.ListenAndServe()
}
