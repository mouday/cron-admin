package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/levigross/grequests"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/enums"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/utils"
	"github.com/robfig/cron"
)

const DATATIME_FORMAT = "2006-01-02 15:04:05"

// 初始化定时任务
func InitCron() {
	db := config.GetDB()
	var list = []model.TaskModel{}

	db.Model(&model.TaskModel{}).Where("status = ?", true).Find(&list)

	for index := range list {
		row := list[index]

		StartTask(row)
	}
}

// 容器
var CronArray sync.Map

type JobParams struct {
	TaskId   string `json:"taskId"`
	Cron     string `json:"cron" `
	Url      string `json:"url" `
	Title    string `json:"title" `
	RunnerId string `json:"runnerId" `
}

func newJob(params JobParams) func() {
	return func() {
		now := time.Now()
		fmt.Println("任务运行：", now.Format(DATATIME_FORMAT))

		item := model.TaskLogModel{}
		item.TaskLogId = utils.GetUuidV4()
		item.TaskId = params.TaskId
		item.Title = params.Title
		item.RunnerId = params.RunnerId
		item.TaskName = "run_job"

		// http://httpbin.org/get
		options := &grequests.RequestOptions{
			JSON: item,
		}

		resp, err := grequests.Post("http://127.0.0.1:5000/api/startTask", options)

		status := enums.TaskStatusStartError
		if err == nil && resp.Ok {
			status = enums.TaskStatusStartSuccess
		}

		item.Status = status

		// database
		db := config.GetDB()
		db.Create(&item)

	}
}

func StopTask(taskId string) {
	cronInstance, ok := CronArray.Load(taskId)

	if ok {
		cronInstance.(*cron.Cron).Stop()
	}

	CronArray.Delete(taskId)
}

func StartTask(row model.TaskModel) {
	// 获取指定cron定时器关闭
	StopTask(row.TaskId)

	params := JobParams{
		TaskId: row.TaskId,
		Cron:   row.Cron,
		Url:    "",
		Title:  row.Title,
	}

	// 每秒执行一次
	cronInstance := cron.New()
	cronInstance.AddFunc(params.Cron, newJob(params))
	cronInstance.Start()

	CronArray.Store(row.TaskId, cronInstance)
}

func ChangeTaskStatus(taskId string, status bool) {
	if status {
		db := config.GetDB()
		row := &model.TaskModel{}
		db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).Find(&row)
		StartTask(*row)
	} else {
		StopTask(taskId)
	}
}
