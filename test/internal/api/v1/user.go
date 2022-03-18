package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/yoga"
	"github.com/xylong/yoga/test/internal/model"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) Register(ctx *gin.Context) string {
	return "注册"
}

func (u *User) Me(ctx *gin.Context) yoga.Model {
	return &model.User{
		ID:       1,
		Phone:    "19999999999",
		Email:    "jingjing@qq.com",
		Unionid:  "",
		Openid1:  "",
		Openid2:  "",
		Avatar:   "",
		Nickname: "静静",
		Password: "123456",
		Birthday: "2020-02-02",
		Gender:   0,
	}
}

func (u *User) Friends(ctx *gin.Context) yoga.Slice {
	users := []*model.User{
		{ID: 1, Nickname: "静静"},
		{ID: 2, Nickname: "JJ"},
	}

	return yoga.MakeModels(users)
}

func (u *User) Logoff(ctx *gin.Context) {

}

func (u *User) Profile(ctx *gin.Context) {

}