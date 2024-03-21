package handler

import (
	"bytes"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// gin拦截response的数据
// https://blog.csdn.net/qq_30505673/article/details/117446285
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// golang(gin)的全局统一异常处理，并统一返回json
// https://blog.csdn.net/u014155085/article/details/106733391
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()

	// blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	// c.Writer = blw
	//加载完 defer recover，继续后续接口调用
	c.Next()

	// params := make(map[string]interface{})

	// // 解析json数据

	// json.Unmarshal(blw.body.Bytes(), &params)

	// fmt.Println("Response body: " + blw.body.String())

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"msg":  "success",
	// 	"data": params,
	// })
	// //没有异常才执行后续接口调用
	// data, exists := c.Get("response_data")
	// if exists {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  "success",
	// 		"data": data,
	// 	})
	// }
	// }
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
