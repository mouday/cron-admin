package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/form"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/service"
	"github.com/mouday/cron-admin/src/vo"
)

func GetTaskLogList(ctx *gin.Context) {
	// taskList := service.GetTaskList()
	// database
	params := &form.PageForm{
		Page:   1,
		Size:   10,
		Status: 0,
	}

	ctx.BindJSON(&params)

	db := config.GetDB()

	taskList := []model.TaskLogModel{}

	var count int64

	if params.Status != 0 {
		db.Model(&model.TaskLogModel{}).Where("status = ?", params.Status).Count(&count)
		db.Model(&model.TaskLogModel{}).Where("status = ?", params.Status).Order("id desc").Limit(params.Size).Offset(params.PageOffset()).Find(&taskList)
	} else {
		db.Model(&model.TaskLogModel{}).Count(&count)
		db.Model(&model.TaskLogModel{}).Order("id desc").Limit(params.Size).Offset(params.PageOffset()).Find(&taskList)

	}

	// ctx.JSON(http.StatusOK, taskList)
	vo.Success(ctx, gin.H{
		"list":  taskList,
		"total": count,
	})
}

type ReportTaskStatusForm struct {
	TaskLogId string `json:"taskLogId"`
	Status    int    `json:"status"`
	Text      string `json:"text"`
}

func ReportTaskStatus(ctx *gin.Context) {
	list := &[]ReportTaskStatusForm{}
	ctx.BindJSON(&list)

	fmt.Println("list", list)

	db := config.GetDB()
	for _, params := range *list {

		db.Model(&model.TaskLogModel{}).Where("task_log_id = ?", params.TaskLogId).Update("status", params.Status)

		// write log
		row := model.TaskLogModel{}
		db.Model(&model.TaskLogModel{}).Where("task_log_id = ?", params.TaskLogId).Find(&row)

		service.AppendLog(row.TaskId, params.TaskLogId, params.Text)
	}

	vo.Success(ctx, nil)
}

func GetTaskLogDetail(ctx *gin.Context) {
	params := &model.TaskLogModel{}

	ctx.BindJSON(&params)

	row := model.TaskLogModel{}

	db := config.GetDB()
	db.Model(&model.TaskLogModel{}).Where("task_log_id = ?", params.TaskLogId).Find(&row)

	content := service.ReadLog(row.TaskId, params.TaskLogId)

	vo.Success(ctx, content)
}
