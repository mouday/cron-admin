package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/vo"
)

func GetTaskLogList(ctx *gin.Context) {
	// taskList := service.GetTaskList()
	// database
	db := config.GetDB()

	taskList := []model.TaskLogModel{}

	db.Model(&model.TaskLogModel{}).Find(&taskList)

	// ctx.JSON(http.StatusOK, taskList)
	vo.Success(ctx, taskList)
}
