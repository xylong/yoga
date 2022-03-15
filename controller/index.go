package controller

import (
	"github.com/gin-gonic/gin"
	"yoga/response"
)

func Index(ctx *gin.Context) *response.Response {
	return response.Resp().Json(gin.H{"data": "hello"})
}
