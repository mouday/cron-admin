package model

import "time"

type TaskModel struct {
	ID         uint
	Title      string
	TaskId     string `gorm:"index"`
	Cron       string
	Url        string
	Running    bool
	CreateTime time.Time `gorm:"type:datetime;autoCreateTime"`
	UpdateTime time.Time `gorm:"type:datetime;autoUpdateTime"`
}

// 自定义表名
func (TaskModel) TableName() string {
	return "tb_task"
}
