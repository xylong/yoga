package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Authorization struct {
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (a *Authorization) Before(ctx *gin.Context) error {
	fmt.Println("auth before", ctx.FullPath())
	return nil
}

func (a *Authorization) After(data interface{}) (interface{}, error) {
	fmt.Println("auth after")
	return nil, nil
}
