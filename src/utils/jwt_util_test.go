package utils

import (
	"fmt"
	"testing"
)

// 获取uuid
func TestCreateToken(t *testing.T) {
	fmt.Println(CreateToken("123", "xxx"))
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjMifQ.0yTSBbJfqdGAmQ3CGHEvHAMBI-AxBvrZd25qBqG6BwQ"
	userId, _ := ParseToken(token, "xxx")
	fmt.Println(userId)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjMifQ.0yTSBbJfqdGAmQ3CGHEvHAMBI-AxBvrZd25qBqG6BwQ"
	fmt.Println(VerifyToken(token, "xxx"))
}
