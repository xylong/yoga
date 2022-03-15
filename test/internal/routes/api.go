package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/yoga"
	"github.com/xylong/yoga/test/internal/api"
	"github.com/xylong/yoga/test/internal/middleware"
)

func Load(engine *gin.Engine) {
	router := yoga.NewRouter(engine)
	router.Group("v1", func(group yoga.Group) {

		group.Registered(yoga.POST, "/register", api.Register)
		group.Registered(yoga.POST, "/login", api.Login)
		group.Registered(yoga.GET, "/users/:id", api.Profile)

		group.Group("", func(group yoga.Group) {
			group.Registered(yoga.GET, "/me", api.Me)
		}, middleware.Authorization)

	}, middleware.Log)
}
