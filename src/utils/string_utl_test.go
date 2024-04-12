package utils

import (
	"fmt"
	"testing"
)

// 获取uuid
func TestGetRandomString(t *testing.T) {
	fmt.Println(GetRandomString(20))
}
