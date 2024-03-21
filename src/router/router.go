package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/api"
)

/* 注册路由 */
func RegistRouter(app *gin.Engine) {

	app.POST("/api/addTask", api.AddTask)
	app.POST("/api/updateTask", api.UpdateTask)
	app.POST("/api/removeTask", api.RemoveTask)
	app.POST("/api/getTask", api.GetTask)
	app.POST("/api/getTaskList", api.GetTaskList)
	app.POST("/api/startTask", api.StartTask)
	app.POST("/api/stopTask", api.StopTask)
}
