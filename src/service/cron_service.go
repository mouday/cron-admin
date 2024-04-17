package service

import (
	"sync"

	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/model"
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

		StartTask(row.TaskId, row.Cron)
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
	Method   string `json:"Method" `
}

func StopTask(taskId string) {
	cronInstance, ok := CronArray.Load(taskId)

	if ok {
		cronInstance.(*cron.Cron).Stop()
	}

	CronArray.Delete(taskId)
}

func StartTask(taskId string, cronExpress string) error {
	// 获取指定cron定时器关闭
	StopTask(taskId)

	// 每秒执行一次
	cronInstance := cron.New()
	err := cronInstance.AddFunc(cronExpress, func() {
		AppendTask(taskId)
	})

	if err != nil {
		return err
	}

	cronInstance.Start()

	CronArray.Store(taskId, cronInstance)

	return nil
}

func ChangeTaskStatus(taskId string, status bool) error {
	var err error

	if status {
		db := config.GetDB()
		row := &model.TaskModel{}
		db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).Find(&row)
		err = StartTask(row.TaskId, row.Cron)
	} else {
		StopTask(taskId)
	}

	return err
}
