package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("content-type", "application/json; charset=utf-8")
		}
	}(),
	)
	r.GET("/", hander)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func hander(ctx *gin.Context) {
	type Hello struct {
		ID   uint64 `json:"id,string"`
		UUID uint64 `json:"uuid"`
	}
	fmt.Println(fmt.Sprintf("%d", uint64(math.MaxUint64)))
	fmt.Println(fmt.Sprintf("%d", uint64(math.MaxUint64-5)))
	fmt.Println(fmt.Sprintf("%d", math.MaxInt64))
	fmt.Println(fmt.Sprintf("%d", math.MaxInt32))
	hello := Hello{ID: math.MaxUint64, UUID: math.MaxUint64}

	JSONOK(ctx, WithItem(&hello))
}

type Res struct {
	Item json.RawMessage `json:"item"`
	Msg  string
}

type ResponseDataOptionFunc func(res *Res)

func WithItem(item interface{}) ResponseDataOptionFunc {
	val := reflect.ValueOf(item)
	// 如果传入的item是一个指针，则判断指针关联的类型是否为结构体
	if val.Kind() == reflect.Struct || (val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct) || val.Kind() == reflect.Map {
		return func(ev *Res) {
			// need to marshal since HTTPData.Item is of type json.RawMessage
			marshalled, err := json.Marshal(item)
			if err != nil {
				ev.Msg = fmt.Sprintf("Failed to marshal item to json: %v", err)
			} else {
				ev.Item = marshalled
			}
		}
	} else {
		return func(ev *Res) {
			ev.Msg = "helloword"
		}
	}
}

// JSONOK custom gin json serialize data
func JSONOK(ctx *gin.Context, opts ...ResponseDataOptionFunc) {
	data := &Res{
		Item: json.RawMessage{},
	}
	for _, op := range opts {
		op(data)
	}
	ctx.JSON(http.StatusOK, data)
}
