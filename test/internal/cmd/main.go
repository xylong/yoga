package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/yoga/test/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Load(r)
	r.Run()
}
