package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/api"
)

/* 注册路由 */
func RegistRouter(app *gin.Engine) {

	// auth
	app.POST("/api/login", api.Login)

	// task
	app.POST("/api/addTask", api.AddTask)
	app.POST("/api/updateTask", api.UpdateTask)
	app.POST("/api/updateTaskStatus", api.UpdateTaskStatus)
	app.POST("/api/removeTask", api.RemoveTask)
	app.POST("/api/getTask", api.GetTask)
	app.POST("/api/getTaskList", api.GetTaskList)
	// app.POST("/api/startTask", api.StartTask)
	// app.POST("/api/stopTask", api.StopTask)

	// log
	app.POST("/api/getTaskLogList", api.GetTaskLogList)
	app.POST("/api/reportTaskStatus", api.ReportTaskStatus)
	app.POST("/api/getTaskLogDetail", api.GetTaskLogDetail)
}
