package main

import (
	"github.com/gin-gonic/gin"
	"yoga/routes"
)

func main() {
	r := gin.Default()
	routes.Load(r)
	r.Run()
}
