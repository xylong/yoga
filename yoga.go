package yoga

import (
	"github.com/gin-gonic/gin"
)

type Yoga struct {
	*gin.Engine
	group *gin.RouterGroup
}

func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}) {
	if h := Convert(handler); h != nil {
		y.group.Handle(httpMethod, relativePath, h)
	}
}

func (y *Yoga) G(path string, callback func())  {
	y.group = y.Group(path)
	callback()
}

func (y *Yoga) Launch() {
	y.Run()
}
