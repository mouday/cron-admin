package handler

import (
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
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
