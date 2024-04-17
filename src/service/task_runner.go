package service

import (
	"fmt"
	"time"

	"github.com/levigross/grequests"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/enums"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/utils"
)

// 启动消费
var ch = make(chan string, 10)

func TaskRunner(taskId string) {

	// query task
	db := config.GetDB()
	taskRow := &model.TaskModel{}
	db.Model(&model.TaskModel{}).Where("task_id = ?", taskId).First(&taskRow)

	fmt.Println("任务运行：", taskRow.Title, taskRow.Url)
	fmt.Println("任务运行开始：", time.Now().Format(DATATIME_FORMAT))

	// start log
	item := model.TaskLogModel{}
	item.TaskLogId = utils.GetUuidV4()
	item.Title = taskRow.Title
	item.TaskId = taskId
	item.Status = enums.TaskStatusStartSuccess

	db.Create(&item)

	// item.Title = params.Title
	// item.RunnerId = params.RunnerId
	// item.TaskName = "run_job"

	// http://httpbin.org/get
	// options := &grequests.RequestOptions{
	// 	Headers: item,
	// }

	resp, err := grequests.DoRegularRequest("GET", taskRow.Url, nil)

	fmt.Println("任务运行结束：", time.Now().Format(DATATIME_FORMAT))

	status := enums.TaskStatusRunError
	if err == nil && resp.Ok {
		status = enums.TaskStatusRunSuccess
	}

	item.Status = status

	// update log
	time := utils.LocalTime{
		Time: time.Now(),
	}

	updateRow := &model.TaskLogModel{}
	updateRow.Status = status
	// updateRow.Result = resp.String()
	updateRow.EndTime = time

	db.Model(&model.TaskLogModel{}).Where("task_log_id = ?", item.TaskLogId).Updates(&updateRow)

	// result
	AppendLog(taskId, item.TaskLogId, resp.String())
}

func Consumer() {
	for {
		taskId, ok := <-ch
		if ok {
			TaskRunner(taskId)
		} else {
			break
		}
	}

	fmt.Println("consumer done")
}

func AppendTask(taskId string) {
	ch <- taskId
}
