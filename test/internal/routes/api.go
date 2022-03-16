package routes

import (
	"github.com/xylong/yoga"
	v1 "github.com/xylong/yoga/test/internal/api/v1"
	"net/http"
)

func Load(yoga *yoga.Yoga) *yoga.Yoga {
	user := v1.NewUser()
	yoga.Handle(http.MethodGet, "register", user.Register)

	return yoga
}
