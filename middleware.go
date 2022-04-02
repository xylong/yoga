package yoga

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware 中间件接口
// 对请求响应进行前置或后置处理
type Middleware interface {
	// Before 执行控制器方法前调用，如获取header信息、参数判断等等
	Before(ctx *gin.Context) error

	// After 控制器方法执行后调用，如封装返回值，日志记录等等
	After(interface{}) (interface{}, error)
}

// middlewares 中间件集合
type middlewares []Middleware

// before 执行所有中间件前置方法
// 前置中间件按先入先出执行
func (m middlewares) before(ctx *gin.Context) {
	for _, f := range m {
		if err := f.Before(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	ctx.Next()
}

// after 执行所有中间件后置方法
// 后置中间件按先入后出执行
func (m middlewares) after(ctx *gin.Context) interface{} {
	data, _ := ctx.Get("return")

	for i := len(m) - 1; i >= 0; i-- {
		if res, err := m[i].After(data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			data = res
		}
	}

	return data
}
