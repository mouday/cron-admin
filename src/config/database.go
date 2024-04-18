package config

import (
	// "gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"fmt"
	"log"
	"os"
	"time"

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
	env := GetEnv()

	var level logger.LogLevel
	if env == "release" {
		level = logger.Error
	} else {
		level = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		// 日志级别
		// Logger: logger.Default.LogMode(logger.Error),
		Logger: newLogger,
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
	username := GetAppAdminUsername()

	userRow := &model.UserModel{}

	db.Model(&model.UserModel{}).Where("username = ?", username).First(&userRow)

	if userRow.Username == "" {
		password := GetAppAdminPassword()

		userRow.UserId = utils.GetUuidV4()
		userRow.Username = username
		userRow.Password = utils.EncodePassword(password)
		userRow.Status = true

		db.Model(&model.UserModel{}).Create(userRow)

		fmt.Println("username:", username)
		fmt.Println("password:", password)
	}
}
