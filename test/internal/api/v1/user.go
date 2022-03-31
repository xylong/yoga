package v1

import (
	"fmt"
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

func (u *User) Login(ctx *gin.Context) string {
	fmt.Println("login")
	return "登录"
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

func (u *User) Logoff(ctx *gin.Context) string {
	fmt.Println("xxoo")
	return "注销"
}

func (u *User) Profile(ctx *gin.Context) string {
	return "用户资料"
}

func Foo(ctx *gin.Context) string {
	return "测试非结构体方法"
}
