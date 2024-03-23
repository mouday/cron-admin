package model

import "github.com/mouday/cron-admin/src/utils"

type TaskModel struct {
	ID         uint            `json:"id"`
	Title      string          `json:"title"`
	TaskId     string          `gorm:"index" json:"taskId"`
	Cron       string          `json:"cron"`
	Url        string          `json:"url"`
	Running    bool            `json:"running"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TaskModel) TableName() string {
	return "tb_task"
}
