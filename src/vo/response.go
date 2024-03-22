package vo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应结果
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
