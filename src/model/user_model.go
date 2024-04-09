package model

import "github.com/mouday/cron-admin/src/utils"

type UserModel struct {
	Id         uint            `json:"-"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (UserModel) TableName() string {
	return "tb_user"
}
