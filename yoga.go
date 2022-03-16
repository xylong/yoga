package yoga

import (
	"github.com/gin-gonic/gin"
)

type Yoga struct {
	*gin.Engine
}

func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}) {
	if h := Convert(handler); h != nil {
		y.Engine.Handle(httpMethod, relativePath, h)
	}
}

func (y *Yoga) Launch() {
	y.Run()
}
