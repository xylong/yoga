package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yoga/context"
	"yoga/controller"
	"yoga/response"
)

type action func(*context.Context) *response.Response

func Load(engine *gin.Engine) {
	engine.GET("/", convert(controller.Index))
}

func convert(f action) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rsp := f(&context.Context{Context: ctx})
		data := rsp.GetData()

		switch item := data.(type) {
		case string:
			ctx.String(http.StatusOK, item)
		case gin.H:
			ctx.JSON(http.StatusOK, item)
		}
	}
}
