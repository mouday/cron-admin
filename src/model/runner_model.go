package model

import "github.com/mouday/cron-admin/src/utils"

type RunnerModel struct {
	Id          uint            `json:"-"`
	RunnerId    string          `json:"runnerId" gorm:"index" `
	Title       string          `json:"title"`
	Url         string          `json:"url"`
	AccessToken string          `json:"accessToken"`
	Status      bool            `json:"status"`
	CreateTime  utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime  utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (RunnerModel) TableName() string {
	return "tb_runner"
}
