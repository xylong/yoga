package yoga

import (
	"github.com/gin-gonic/gin"
)

type Yoga struct {
	*gin.Engine
	Path string
}

func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

func (y *Yoga) Group(path string, callback func(Group)) {
	callback(Group{
		Engine: y.Engine,
		Path:   path,
	})
}

func (y *Yoga) Launch() {
	y.Run()
}
