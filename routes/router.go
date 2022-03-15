package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yoga/context"
	"yoga/response"
)

const (
	GET    method = 0x000000
	POST   method = 0x000001
	PUT    method = 0x000002
	DELETE method = 0x000003
	ANY    method = 0x000004
)

type (
	method int
	action func(*context.Context) *response.Response
)

type router struct {
	*gin.Engine
}

func newRouter(engine *gin.Engine) *router {
	return &router{Engine: engine}
}

func (r *router) Group(path string, callback func(group), middlewares ...context.HandlerFunc) {
	callback(group{
		engine:      r.Engine,
		path:        path,
		middlewares: middlewares,
	})
}

type group struct {
	engine      *gin.Engine
	path        string
	middlewares []context.HandlerFunc
}

func (g group) Group(path string, callback func(group), middlewares ...context.HandlerFunc) {
	// 父级的中间件在前面
	g.middlewares = append(g.middlewares, middlewares...)
	g.path += path
	callback(g)
}

func (g group) Registered(m method, url string, a action, middlewares ...context.HandlerFunc) {
	// 父类中间件+当前action中间件+action
	handlers := make([]gin.HandlerFunc, len(g.middlewares)+len(middlewares)+1)

	// 添加当前action的中间件
	g.middlewares = append(g.middlewares, middlewares...)

	// 将中间件转换为gin.HandlerFunc
	for key, middleware := range g.middlewares {
		temp := middleware
		handlers[key] = func(c *gin.Context) {
			temp(&context.Context{Context: c})
		}
	}
	//添加action
	handlers[len(g.middlewares)] = convert(a)
	finalUrl := g.path + url

	switch m {
	case GET:
		g.engine.GET(finalUrl, handlers...)
	case POST:
		g.engine.GET(finalUrl, handlers...)
	case PUT:
		g.engine.PUT(finalUrl, handlers...)
	case DELETE:
		g.engine.DELETE(finalUrl, handlers...)
	case ANY:
		g.engine.Any(finalUrl, handlers...)
	}
}

func convert(f action) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rsp := f(&context.Context{Context: ctx})
		data := rsp.GetData()

		switch item := data.(type) {
		case string:
			ctx.String(http.StatusOK, item)
		case gin.H:
			ctx.JSON(http.StatusOK, item)
		}
	}
}
