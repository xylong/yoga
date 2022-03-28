package yoga

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Yoga struct {
	*gin.Engine
	group       string
	middlewares []Middleware
}

// Ignite 初始化
func Ignite() *Yoga {
	yoga := &Yoga{Engine: gin.New()}
	yoga.Use(gin.Logger(), gin.Recovery())

	return yoga
}

// Group 路由分组
// callback中传入复制yoga的非指针值，达到在函数作用域中隔离分组设置
func (y *Yoga) Group(group string, callback func(Yoga), middlewares ...Middleware) {
	g := *y
	// 加入父级中间件
	g.joinMiddleware(middlewares...)
	g.group += fmt.Sprintf("/%s", strings.Trim(group, "/"))
	callback(g)
}

// Handle 重载gin的Handle方法
// 对控制器方法进行签名判断
// 如string、slice、json等返回类型的函数签名
func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}, middlewares ...Middleware) {
	if h := Convert(handler); h != nil {
		// 加入路由级中间件
		y.joinMiddleware(middlewares...)
		finalUrl := y.group + "/" + relativePath

		// 执行顺讯：全局中间件->父级中间件->路由级中间件->路由回调
		y.Engine.Handle(httpMethod, finalUrl, func(context *gin.Context) {
			context.Set("middlewares", y.middlewares)

			for _, middleware := range y.middlewares {
				if err := middleware.Before(context); err != nil {
					context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
				}

				context.Next()
			}
		}, h)
	}
}

// joinMiddleware 加入中间件
// 中间件顺序：全局中间件->父级中间件->路由级中间件
func (y *Yoga) joinMiddleware(middlewares ...Middleware) {
	y.middlewares = append(y.middlewares, middlewares...)
}

// Launch 启动服务
func (y *Yoga) Launch() {
	y.Run()
}
