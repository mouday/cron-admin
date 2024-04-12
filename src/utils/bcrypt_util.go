package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func EncodePassword(password string) string {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hash)

}

// 验证密码
func VerifyPassword(plainPwd string, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err == nil
}
