package model

import "github.com/mouday/cron-admin/src/utils"

type TaskLogModel struct {
	ID         uint            `json:"id"`
	TaskLogId  string          `gorm:"index" json:"taskLogId"`
	TaskId     string          `json:"taskId"`
	Status     int             `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TaskLogModel) TableName() string {
	return "tb_log_task"
}
