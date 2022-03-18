package yoga

import (
	"github.com/gin-gonic/gin"
)

type Yoga struct {
	*gin.Engine
	group string
}

func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}) {
	if h := Convert(handler); h != nil {
		url := y.group + "/" + relativePath
		y.Engine.Handle(httpMethod, url, h)
	}
}

func (y *Yoga) G(path string, callback func(Yoga)) {
	y.group += path
	callback(*y)
}

func (y *Yoga) Launch() {
	y.Run()
}
