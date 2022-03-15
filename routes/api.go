package routes

import (
	"github.com/gin-gonic/gin"
	"yoga/controller"
)

func Load(engine *gin.Engine) {
	router := newRouter(engine)
	router.Group("v1", func(g group) {

		g.Registered(POST, "/register", controller.Register)
		g.Registered(POST, "/login", controller.Login)

		g.Group("/users", func(g group) {
			g.Registered(GET, "/:id", controller.Profile)
		})
	})
	engine.GET("/", convert(controller.Index))
}
