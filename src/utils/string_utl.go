package utils

import (
	"github.com/pochard/commons/randstr"
)

// letters + numbers + other visible ascii chars
func GetRandomString(n int) string {
	return randstr.RandomAscii(n)
}
