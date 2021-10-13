package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	res := []string{"http://192.168.199.109"}
	bts, _ := json.Marshal(res)
	fmt.Println(string(bts))

}

func hander(ctx *gin.Context) {
	action := ctx.Param("action")

	ctx.JSON(200, action)
}
