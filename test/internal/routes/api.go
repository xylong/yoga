package routes

import (
	"github.com/xylong/yoga"
	v1 "github.com/xylong/yoga/test/internal/api/v1"
	"github.com/xylong/yoga/test/internal/middleware"
	"net/http"
)

func Load(y *yoga.Yoga) *yoga.Yoga {
	y.Group("v1", func(y yoga.Yoga) {
		user := v1.NewUser()
		y.Handle(http.MethodGet, "register", user.Register)

		y.Group("/a", func(y yoga.Yoga) {
			y.Handle(http.MethodGet, "me", user.Me)
			y.Group("/b", func(y yoga.Yoga) {
				y.Handle(http.MethodGet, "friends", user.Friends)

				y.Group("c", func(y yoga.Yoga) {
					y.Handle(http.MethodPost, "users/:id", user.Profile)
				}, middleware.NewAuthorization())
			})
		}, middleware.NewCsrf())
		y.Handle(http.MethodDelete, "logoff", user.Logoff, middleware.NewAuthorization())
		y.Handle(http.MethodGet, "foo", v1.Foo)
	})

	return y
}
