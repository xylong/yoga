package yoga

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Yoga struct {
	*gin.Engine
}

func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}) {
	if h, ok := handler.(func(ctx *gin.Context) string); ok {
		y.Engine.Handle(httpMethod, relativePath, func(context *gin.Context) {
			context.String(http.StatusOK, h(context))
		})
	}
}

func (y *Yoga) Launch() {
	y.Run()
}
