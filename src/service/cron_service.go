package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/utils"
	"github.com/robfig/cron"
)

const DATATIME_FORMAT = "2006-01-02 15:04:05"

// 初始化定时任务
func InitCron() {
	db := config.GetDB()
	var list = []model.TaskModel{}

	db.Model(&model.TaskModel{}).Where("running = ?", true).Find(&list)

	for index := range list {
		taskModel := list[index]
		params := &JobParams{}

		params.Cron = taskModel.Cron
		params.TaskId = taskModel.TaskId
		params.Url = taskModel.Url

		AddTask(params)
	}
}

// 容器
var CronArray sync.Map

type JobParams struct {
	TaskId  string `json:"taskId"`
	Cron    string `json:"cron" `
	Url     string `json:"url" `
	Running bool   `json:"running" `
}

type Job struct {
	Params *JobParams
	Cron   *cron.Cron
}

func task() {
	now := time.Now()
	fmt.Println("任务运行：", now.Format(DATATIME_FORMAT))
}

func job(params *JobParams) func() {
	return func() {
		now := time.Now()
		fmt.Println("任务运行：", now.Format(DATATIME_FORMAT))

		// "http://httpbin.org/get"
		// resp, _ := grequests.Get(params.Url, nil)

		// if resp.Ok {
		// 	fmt.Println(resp.String())
		// } else {
		// 	fmt.Println(resp.Error)
		// }
	}
}

func AddTask(params *JobParams) *JobParams {
	if params.TaskId == "" {
		params.TaskId = utils.GetUuid()
	} else {
		oldParams := GetTask(params.TaskId)

		if params.Url == "" {
			params.Url = oldParams.Url
		}

		if params.Cron == "" {
			params.Cron = oldParams.Cron
		}

		RemoveTask(params.TaskId)
	}

	// 每秒执行一次
	c := cron.New()
	c.AddFunc(params.Cron, job(params))
	c.Start()

	params.Running = true

	job := Job{
		Params: params,
		Cron:   c,
	}

	CronArray.Store(params.TaskId, job)

	// database
	db := config.GetDB()

	item := model.TaskModel{}
	item.Cron = params.Cron
	item.Url = params.Url
	item.TaskId = params.TaskId
	item.Running = params.Running

	db.Create(&item)

	return params
}

func RemoveTask(taskId string) {
	// 获取指定cron定时器关闭
	StopTask(taskId)
	CronArray.Delete(taskId)

	// database
	db := config.GetDB()

	task := model.TaskModel{}

	db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).Delete(&task)
}

func StopTask(taskId string) {
	// 获取指定cron定时器关闭

	task, ok := CronArray.Load(taskId)
	if ok {
		job := task.(Job)
		job.Cron.Stop()
		job.Params.Running = false
	}

	CronArray.Store(taskId, task)

	// database
	db := config.GetDB()
	taskModel := model.TaskModel{}
	taskModel.Running = false
	db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).Updates(&taskModel)
}

func StartTask(taskId string) {
	// 获取指定cron定时器关闭

	task, ok := CronArray.Load(taskId)
	if ok {

		job := task.(Job)
		job.Cron.Start()

		job.Params.Running = true
	}

	CronArray.Store(taskId, task)

	// database
	db := config.GetDB()
	taskModel := model.TaskModel{}
	taskModel.Running = true
	db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).Updates(&taskModel)
}

func GetTask(taskId string) *JobParams {
	task, ok := CronArray.Load(taskId)

	if ok {
		return task.(Job).Params
	} else {
		return &JobParams{}
	}
}

func GetTaskList() []*JobParams {
	taskList := []*JobParams{}
	CronArray.Range(func(key, value interface{}) bool {
		job := value.(Job)
		taskList = append(taskList, job.Params)
		return true
	})

	return taskList
}
