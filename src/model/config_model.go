package model

import "github.com/mouday/cron-admin/src/utils"

// 配置
type ConfigModel struct {
	Id         uint            `json:"-"`
	Key        string          `json:"key"  gorm:"index"`
	Value      string          `json:"value"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (ConfigModel) TableName() string {
	return "tb_config"
}
