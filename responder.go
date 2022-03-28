package yoga

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

func init() {
	ResponderList = []Responder{
		new(NilResponder),
		new(StringResponder),
		new(JsonResponder),
		new(SliceResponder),
		new(ViewResponder),
	}
}

var (
	// ResponderList 响应类型列表
	ResponderList []Responder
)

// Responder 控制器返回值
type Responder interface {
	Return() gin.HandlerFunc
}

// NilResponder 无返回值响应
type NilResponder func(*gin.Context)

func (r NilResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		r(context)
	}
}

// StringResponder 字符串响应
type StringResponder func(*gin.Context) string

func (r StringResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		data := r(context)
		if middlewares, exists := context.Get("middlewares"); exists {
			for _, middleware := range middlewares.([]Middleware) {
				middleware.After(data)
			}
		}

		context.String(http.StatusOK, data)
	}
}

// JsonResponder 实体响应
type JsonResponder func(*gin.Context) Model

func (r JsonResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, r(context))
	}
}

// SliceResponder 切片响应
type SliceResponder func(*gin.Context) Slice

func (r SliceResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		if _, err := context.Writer.WriteString(string(r(context))); err != nil {
			log.Println(err)
		}
	}
}

type (
	// View 视图模板
	View string

	// ViewResponder 视图响应
	ViewResponder func(*gin.Context) View
)

func (r ViewResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, fmt.Sprintf("%s.html", string(r(context))), context.Keys)
	}
}

// Convert 将方法转换为控制器响应方法
// 因为无法将interface断言成Responder或StringResponder，
// 所以需要通过反射来转换
func Convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range ResponderList {
		// 因为ResponderList存放的都是类型的初始化指针，所以要用Elem
		val := reflect.ValueOf(responder).Elem()
		if value.Type().ConvertibleTo(val.Type()) {
			// 将handler的反射对象的值赋值给val
			// val此时是反射对象
			val.Set(value)
			// Interface获取val的值,
			// 将值断言为Responder来调用Handle
			return val.Interface().(Responder).Return()
		}
	}

	return nil
}
