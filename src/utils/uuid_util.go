package utils

import (
	"crypto/md5"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

// 获取uuid
func GetUuidV4() string {
	return uuid.NewV4().String()
}

// 获取uuid
func GetUuid() string {
	u := uuid.NewV4()
	return Get16MD5Encode(u.String())
}

// 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
