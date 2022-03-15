package middleware

import (
	"fmt"
	"github.com/xylong/yoga"
)

func Authorization(*yoga.Context) {
	fmt.Println("auth")
}

func Log(*yoga.Context) {
	fmt.Println("log")
}
