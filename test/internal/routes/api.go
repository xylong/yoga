package routes

import (
	"github.com/xylong/yoga"
	v1 "github.com/xylong/yoga/test/internal/api/v1"
	"net/http"
)

func Load(y *yoga.Yoga) *yoga.Yoga {
	y.Group("v1", func(group yoga.Group) {
		user := v1.NewUser()
		group.Handle(http.MethodGet, "register", user.Register)

		group.Group("/a", func(group yoga.Group) {
			group.Handle(http.MethodGet, "me", user.Me)

			group.Group("/b", func(group yoga.Group) {
				group.Handle(http.MethodGet, "friends", user.Friends)

				group.Group("/c", func(group yoga.Group) {
					group.Handle(http.MethodPost, "/users/:id", user.Profile)
				})
			})
		})

		group.Handle(http.MethodDelete, "logoff", user.Logoff)
	})

	return y
}
