package yoga

import (
	"encoding/json"
	"log"
)

// Model 对象实例
type Model interface {
	String() string
}

// Slice 实体切片
type Slice string

// MakeModels 创建实体切片
func MakeModels(value interface{}) Slice {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
	}

	return Slice(bytes)
}
