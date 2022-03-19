package yoga

import "github.com/gin-gonic/gin"

// Middleware 中间件接口
// 对请求响应进行前置或后置处理
type Middleware interface {
	// Before 执行控制器方法前调用，如获取header信息、参数判断等等
	Before(ctx *gin.Context) error

	// After 控制器方法执行后调用，如封装返回值，日志记录等等
	After(interface{}) (interface{}, error)
}
