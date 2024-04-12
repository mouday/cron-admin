package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/handler"
	"github.com/mouday/cron-admin/src/router"
	"github.com/mouday/cron-admin/src/service"
)

//go:embed public/*
var local embed.FS

func main() {
	// app
	app := gin.New()

	// 全局异常捕获
	app.Use(handler.Recover)

	app.Use(handler.AuthMiddleware())

	// 注册路由
	router.RegistRouter(app)

	// 数据库迁移
	config.Migrate()

	// 初始化数据
	config.InitData()

	// 初始化定时任务
	service.InitCron()

	// 【Go语言】gin + go:embed 打包静态资源文件
	// ref: https://blog.csdn.net/Regulations/article/details/128858670
	fp, _ := fs.Sub(local, "public")
	app.StaticFS("/", http.FS(fp))

	// 监听并在 http://127.0.0.1:8082 上启动服务
	err := app.Run(":8082")

	if err != nil {
		fmt.Println(err)
	}
}
