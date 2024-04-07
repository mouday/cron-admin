package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/form"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/service"
	"github.com/mouday/cron-admin/src/utils"
	"github.com/mouday/cron-admin/src/vo"
)

type TaskForm struct {
	TaskId string `json:"taskId"`
	Title  string `json:"title"`
	Cron   string `json:"cron" `
	Url    string `json:"url" `
	Status bool   `json:"status" `
}

func AddTask(ctx *gin.Context) {
	form := &TaskForm{}
	ctx.BindJSON(&form)

	row := &model.TaskModel{
		Title:  form.Title,
		Cron:   form.Cron,
		Url:    form.Url,
		Status: form.Status,
		TaskId: utils.GetUuidV4(),
	}

	db := config.GetDB()
	db.Model(&model.TaskModel{}).Create(&row)

	service.ChangeTaskStatus(row.TaskId, row.Status)

	vo.Success(ctx, nil)
}

func UpdateTask(ctx *gin.Context) {
	row := &model.TaskModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.TaskModel{}).Where("task_id = ?", row.TaskId).Updates(&row)

	service.ChangeTaskStatus(row.TaskId, row.Status)

	vo.Success(ctx, nil)
}

func UpdateTaskStatus(ctx *gin.Context) {
	params := &model.TaskModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.TaskModel{}).Where("task_id = ?", params.TaskId).Update("status", params.Status)

	service.ChangeTaskStatus(params.TaskId, params.Status)

	vo.Success(ctx, nil)
}

func RemoveTask(ctx *gin.Context) {
	row := &model.TaskModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("task_id = ?", row.TaskId).Delete(&model.TaskModel{})

	// service.RemoveTask(params.TaskId)
	// ctx.String(http.StatusOK, "hello")
	vo.Success(ctx, nil)
}

func GetTask(ctx *gin.Context) {
	// params := &service.JobParams{}

	// // 解析json数据
	// rawData, _ := ctx.GetRawData()

	// json.Unmarshal(rawData, &params)

	// task := service.GetTask(params.TaskId)
	row := &model.TaskModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Model(&model.TaskModel{}).Where("task_id = ?", row.TaskId).Find(&row)

	// ctx.JSON(http.StatusOK, task)
	vo.Success(ctx, row)
}

func GetTaskList(ctx *gin.Context) {

	params := &form.PageForm{
		Page:   1,
		Size:   10,
		Status: 0,
	}

	ctx.BindJSON(&params)

	db := config.GetDB()

	taskList := []model.TaskModel{}
	var count int64

	db.Model(&model.TaskModel{}).Count(&count)

	db.Model(&model.TaskModel{}).Order("id desc").Limit(params.Size).Offset(params.PageOffset()).Find(&taskList)

	vo.Success(ctx, gin.H{
		"list":  taskList,
		"total": count,
	})
}
