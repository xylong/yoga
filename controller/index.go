package controller

import (
	"yoga/context"
	"yoga/response"
)

func Index(ctx *context.Context) *response.Response {
	return response.Resp().String(ctx.Domain())
}
