package utils

import (
	"fmt"
	"testing"
)

// 获取uuid
func TestEncodePassword(t *testing.T) {
	fmt.Println(EncodePassword("123456"))
}

func TestVerifyPassword(t *testing.T) {
	fmt.Println(VerifyPassword("123456", "$2a$04$IQ7OPfxRyA4nhDEVi5jhGuAbpog1oy6gfgpIXfPyWdp4sGZc7s02q"))
}
