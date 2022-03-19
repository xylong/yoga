package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Csrf struct {
}

func NewCsrf() *Csrf {
	return &Csrf{}
}

func (c *Csrf) Before(ctx *gin.Context) error {
	fmt.Println("csrf Before", ctx.FullPath())
	return nil
}

func (c *Csrf) After(data interface{}) (interface{}, error) {
	fmt.Println("csrf after")
	return nil, nil
}
