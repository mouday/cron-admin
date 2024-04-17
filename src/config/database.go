package config

import (
	// "gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局秘钥
var SCERET string

// 获取数据库连接
func GetDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		// 日志级别
		// Logger: logger.Default.LogMode(logger.Error),
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db
}

// 迁移数据库
func Migrate() {
	db := GetDB()
	db.AutoMigrate(
		&model.TaskModel{},
		&model.TaskLogModel{},
		&model.UserModel{},
		&model.ConfigModel{},
		&model.RunnerModel{},
	)
}

func InitData() {
	InitConfigData()
	InitUserData()
}

func InitConfigData() {
	db := GetDB()

	configRow := &model.ConfigModel{}
	db.Model(&model.ConfigModel{}).Where("key = ?", "secret").First(&configRow)

	if configRow.Value == "" {
		SCERET = utils.GetRandomString(20)

		configRow.Key = "secret"
		configRow.Value = SCERET

		db.Model(&model.ConfigModel{}).Create(configRow)
	} else {
		SCERET = configRow.Value
	}
}

func InitUserData() {
	db := GetDB()

	userRow := &model.UserModel{}
	db.Model(&model.UserModel{}).Where("username = ?", "admin").First(&userRow)

	if userRow.Username == "" {
		userRow.UserId = utils.GetUuidV4()
		userRow.Username = "admin"
		userRow.Password = utils.EncodePassword("123456")
		userRow.Status = true

		db.Model(&model.UserModel{}).Create(userRow)
	}
}
