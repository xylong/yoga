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

func (r *router) Group(path string, callback func(group)) {
	callback(group{
		engine: r.Engine,
		path:   path,
	})
}

type group struct {
	engine *gin.Engine
	path   string
}

func (g group) Group(path string, callback func(group)) {
	g.path += path
	callback(g)
}

func (g group) Registered(m method, url string, a action) {
	handlerFunc := convert(a)
	finalUrl := g.path + url

	switch m {
	case GET:
		g.engine.GET(finalUrl, handlerFunc)
	case POST:
		g.engine.GET(finalUrl, handlerFunc)
	case PUT:
		g.engine.PUT(finalUrl, handlerFunc)
	case DELETE:
		g.engine.DELETE(finalUrl, handlerFunc)
	case ANY:
		g.engine.Any(finalUrl, handlerFunc)
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
