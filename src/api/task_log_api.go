package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/vo"
)

type PageForm struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	Status int `json:"status"`
}

func (pageForm PageForm) PageOffset() int {
	return (pageForm.Page - 1) * pageForm.Size
}

func GetTaskLogList(ctx *gin.Context) {
	// taskList := service.GetTaskList()
	// database
	params := &PageForm{
		Page:   1,
		Size:   10,
		Status: 0,
	}

	ctx.BindJSON(&params)

	db := config.GetDB()

	taskList := []model.TaskLogModel{}

	db.Model(&model.TaskLogModel{}).Find(&taskList).Limit(params.Size).Offset(params.PageOffset())

	// ctx.JSON(http.StatusOK, taskList)
	vo.Success(ctx, taskList)
}
