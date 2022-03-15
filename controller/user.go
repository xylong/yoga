package controller

import (
	"github.com/gin-gonic/gin"
	"yoga/context"
	"yoga/response"
)

func Register(ctx *context.Context) *response.Response {
	return response.Resp().String("注册")
}

func Login(ctx *context.Context) *response.Response {
	return response.Resp().String("登录")
}

func Profile(ctx *context.Context) *response.Response {
	return response.Resp().Json(gin.H{"data": ctx.Param("id")})
}

func Me(ctx *context.Context) *response.Response {
	return response.Resp().Json(gin.H{"data": "user"})
}
