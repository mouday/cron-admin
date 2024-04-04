package model

import "github.com/mouday/cron-admin/src/utils"

type TaskModel struct {
	Id         uint            `json:"-"`
	Title      string          `json:"title"`
	Cron       string          `json:"cron"`
	Url        string          `json:"url"`
	Status     bool            `json:"status"`
	TaskId     string          `gorm:"index" json:"taskId"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TaskModel) TableName() string {
	return "tb_task"
}
