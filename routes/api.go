package routes

import (
	"github.com/gin-gonic/gin"
	"yoga/controller"
	"yoga/middleware"
)

func Load(engine *gin.Engine) {
	router := newRouter(engine)
	router.Group("v1", func(g group) {

		g.Registered(POST, "/register", controller.Register)
		g.Registered(POST, "/login", controller.Login)
		g.Registered(GET, "/users/:id", controller.Profile)

		g.Group("", func(g group) {
			g.Registered(GET, "/me", controller.Me)
		}, middleware.Authorization)

	}, middleware.Log)
}
