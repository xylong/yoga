package yoga

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Yoga struct {
	*gin.Engine
	group string
}

// Ignite 初始化
func Ignite() *Yoga {
	return &Yoga{Engine: gin.New()}
}

// Group 路由分组
// callback中传入复制yoga的非指针值，达到在函数作用域中隔离分组设置
func (y *Yoga) Group(group string, callback func(Yoga)) {
	g := *y
	g.group += fmt.Sprintf("/%s", strings.Trim(group, "/"))
	callback(g)
}

// Handle 重载gin的Handle方法
// 对控制器方法进行签名判断
// 如string、slice、model等返回类型的函数签名
//
// 如果转换成功，则将控制器方法注册到路由
func (y *Yoga) Handle(httpMethod, relativePath string, handler interface{}) {
	if h := Convert(handler); h != nil {
		url := y.group + "/" + relativePath
		y.Engine.Handle(httpMethod, url, h)
	}
}

// Launch 启动服务
func (y *Yoga) Launch() {
	y.Run()
}
