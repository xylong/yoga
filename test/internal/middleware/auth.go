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
	fmt.Println("auth before")
	return nil
}

func (a *Authorization) After(data interface{}) (interface{}, error) {
	fmt.Println("auth after")
	return gin.H{
		"code": 0,
		"data": data,
	}, nil
}
