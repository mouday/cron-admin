package config

import (
	"os"

	"github.com/mouday/cron-admin/src/utils"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func GetEnv() string {
	env := os.Getenv("GIN_MODE")

	if env == "" {
		// env = "release"
		env = "release"
	}

	return env
}
func GetAppRunAddress() string {
	// 启动服务
	appRunAddress := os.Getenv("APP_RUN_ADDRESS")

	if appRunAddress == "" {
		appRunAddress = "127.0.0.1:8000"
	}

	return appRunAddress
}

func GetAppAdminUsername() string {
	username := os.Getenv("APP_ADMIN_USERNAME")

	if username == "" {
		username = "admin"
	}

	return username

}

func GetAppAdminPassword() string {
	password := os.Getenv("APP_ADMIN_PASSWORD")

	if password == "" {
		password = utils.GetRandomString(10)
	}

	return password

}
