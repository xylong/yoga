package yoga

import "github.com/gin-gonic/gin"

type Group struct {
	*gin.Engine
	Path string
}

func (g Group) Group(path string, callback func(Group)) {
	g.Path += path
	callback(g)
}

func (g Group) Handle(httpMethod, relativePath string, handler interface{}) {
	if h := Convert(handler); h != nil {
		url := g.Path + "/" + relativePath
		g.Engine.Handle(httpMethod, url, h)
	}
}
