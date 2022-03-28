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

		y.Group("", func(y yoga.Yoga) {
			y.Handle(http.MethodGet, "me", user.Me)
			y.Handle(http.MethodGet, "friends", user.Friends)
			y.Handle(http.MethodDelete, "logoff", user.Logoff)
		}, middleware.NewAuthorization())

		y.Handle(http.MethodPost, "register", user.Register)
		y.Handle(http.MethodPost, "login", user.Login)
	})

	return y
}
