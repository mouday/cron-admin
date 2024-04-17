package handler

import (
	"fmt"
	"log"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/utils"
	"github.com/mouday/cron-admin/src/vo"
)

// gin拦截response的数据
// https://blog.csdn.net/qq_30505673/article/details/117446285

// golang(gin)的全局统一异常处理，并统一返回json
// https://blog.csdn.net/u014155085/article/details/106733391
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			vo.Error(c, 1, errorToString(r))

			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()

	c.Next()

}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

var ALLOW_URLS = []string{
	"/api/login",

	// 静态文件放行
	"/",
	".*\\.html",
	".*\\.js",
	".*\\.css",
	".*\\.svg",
	".*\\.ico",
	".*\\.gif",
}

func InArray(url string, urls []string) bool {

	for _, v := range urls {
		if strings.Contains(v, "*") {
			matched, _ := regexp.MatchString(v, url)
			if matched {
				return true
			}
		} else if v == url {
			return true
		}
	}

	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		url := c.Request.RequestURI
		fmt.Println("url:", c.Request.RequestURI)

		if !InArray(url, ALLOW_URLS) {
			token := c.Request.Header.Get("X-Token")
			fmt.Println("token:", token)

			if token == "" || !utils.VerifyToken(token, config.SCERET) {
				vo.Error(c, 403, "Token验证失败")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
