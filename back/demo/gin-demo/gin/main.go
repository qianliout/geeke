package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", hander)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func hander(ctx *gin.Context) {
	action := ctx.Query("action")
	fmt.Println("action is ", action)
	ctx.JSON(200, action)
}
