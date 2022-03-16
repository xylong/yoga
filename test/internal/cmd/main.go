package main

import (
	"github.com/xylong/yoga"
	"github.com/xylong/yoga/test/internal/routes"
)

func main() {
	routes.Load(yoga.Ignite()).Launch()
}
