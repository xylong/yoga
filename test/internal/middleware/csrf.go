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
	fmt.Println("csrf Before")
	return nil
}

func (c *Csrf) After(data interface{}) (interface{}, error) {
	fmt.Println("csrf after")
	return data, nil
}
