package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/handler"
	"github.com/mouday/cron-admin/src/router"
	"github.com/mouday/cron-admin/src/service"
)

func main() {
	// app
	app := gin.New()

	// 全局异常捕获
	app.Use(handler.Recover)

	// 注册路由
	router.RegistRouter(app)

	//
	config.Migrate()

	service.InitCron()

	// 监听并在127.0.0.1:8082 上启动服务
	err := app.Run(":8082")

	if err != nil {
		fmt.Println(err)
	}
}
