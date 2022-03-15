package middleware

import (
	"fmt"
	"yoga/context"
)

func Authorization(*context.Context) {
	fmt.Println("auth")
}

func Log(*context.Context) {
	fmt.Println("log")
}
