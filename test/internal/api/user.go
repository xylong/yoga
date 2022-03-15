package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/yoga"
)

func Register(ctx *yoga.Context) *yoga.Response {
	return yoga.Resp().String("注册")
}

func Login(ctx *yoga.Context) *yoga.Response {
	return yoga.Resp().String("登录")
}

func Profile(ctx *yoga.Context) *yoga.Response {
	return yoga.Resp().Json(gin.H{"data": ctx.Param("id")})
}

func Me(ctx *yoga.Context) *yoga.Response {
	return yoga.Resp().Json(gin.H{"data": "user"})
}
