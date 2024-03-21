package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/service"
	"github.com/mouday/cron-admin/src/vo"
)

func AddTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	// rawData, _ := ctx.GetRawData()

	// json.Unmarshal(rawData, &params)
	ctx.BindJSON(&params)

	fmt.Println(params)

	jobParams := service.AddTask(params)

	// ctx.JSON(http.StatusOK, jobParams)
	vo.Success(ctx, jobParams)
}

func UpdateTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	rawData, _ := ctx.GetRawData()

	json.Unmarshal(rawData, &params)
	fmt.Println(params)

	jobParams := service.AddTask(params)

	// ctx.JSON(http.StatusOK, jobParams)
	vo.Success(ctx, jobParams)
}

func RemoveTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	rawData, _ := ctx.GetRawData()

	json.Unmarshal(rawData, &params)

	service.RemoveTask(params.TaskId)
	// ctx.String(http.StatusOK, "hello")
	vo.Success(ctx, nil)
}

func GetTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	rawData, _ := ctx.GetRawData()

	json.Unmarshal(rawData, &params)

	task := service.GetTask(params.TaskId)
	// ctx.JSON(http.StatusOK, task)
	vo.Success(ctx, task)
}

func GetTaskList(ctx *gin.Context) {
	taskList := service.GetTaskList()
	// ctx.JSON(http.StatusOK, taskList)
	vo.Success(ctx, taskList)
}

func StartTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	rawData, _ := ctx.GetRawData()

	json.Unmarshal(rawData, &params)

	service.StartTask(params.TaskId)
	// ctx.JSON(http.StatusOK, "ok")
	vo.Success(ctx, nil)
}

func StopTask(ctx *gin.Context) {
	params := &service.JobParams{}

	// 解析json数据
	rawData, _ := ctx.GetRawData()

	json.Unmarshal(rawData, &params)

	service.StopTask(params.TaskId)
	// ctx.JSON(http.StatusOK, "ok")
	vo.Success(ctx, nil)
}
