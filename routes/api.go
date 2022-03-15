package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yoga/controller"
	"yoga/response"
)

func Load(engine *gin.Engine) {
	engine.GET("/", convert(controller.Index))
}

func convert(f func(ctx *gin.Context) *response.Response) gin.HandlerFunc {
	return func(context *gin.Context) {
		rsp := f(context)
		data := rsp.GetData()

		switch item := data.(type) {
		case string:
			context.String(http.StatusOK, item)
		case gin.H:
			context.JSON(http.StatusOK, item)
		}
	}
}
