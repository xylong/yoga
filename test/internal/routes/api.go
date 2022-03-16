package routes

import (
	"github.com/xylong/yoga"
	v1 "github.com/xylong/yoga/test/internal/api/v1"
	"net/http"
)

func Load(yoga *yoga.Yoga) *yoga.Yoga {
	user := v1.NewUser()
	yoga.Handle(http.MethodGet, "register", user.Register)
	yoga.Handle(http.MethodGet, "me", user.Me)
	yoga.Handle(http.MethodGet, "friends", user.Friends)

	return yoga
}
