package config

import "os"

func GetEnv() string {
	env := os.Getenv("GIN_MODE")

	if env == "" {
		// env = "release"
		env = "debug"
	}

	return env
}
