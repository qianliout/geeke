package main

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	c := Conn{
		msgChan: msgCh(),
		fuChan:  make(chan func(string2 string)),
	}
	go func() {
		c.fuChan <- fun(1)
		c.fuChan <- fun(2)
		c.fuChan <- fun(3)
	}()

	go c.Hello()
	// r := gin.Default()
	// r.GET("/ws/scanner", handler)
	// r.GET("/hello", hello)
	// r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	time.Sleep(10 * time.Second)
}

type Conn struct {
	msgChan chan string
	fuChan  chan func(string2 string)

	// fn1  chan func(string)
	// fn2  chan func(string)
}

func msgCh() chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			out <- "hello"
			time.Sleep(time.Second)
			out <- "word"
		}
	}()
	return out
}
func fun(i int) func(string2 string) {
	return func(s string) {
		fmt.Println("func ", i, s)
	}
}

func (s Conn) Hello() {
	for {
		msg := <-s.msgChan

		fus := make([]func(string), 0)

		if len(s.fuChan) == 0 {
			fmt.Println("no cl")
			continue
		}
		for i := range s.fuChan {
			s.fuChan[i](msg)
			fus = append(fus, fun(i+10))
		}
		s.fuChan = fus
		go func() { s.msgChan <- "fuck" }()
	}
}

func handler(ctx *gin.Context) {
	client, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Println("error ", err.Error())
		return
	}
	id := time.Now().Unix()

	for {
		fmt.Println("send one data")
		type Hello struct {
			ID   int64  `json:"id,string"`
			UUID string `json:"uuid"`
		}
		data := Hello{ID: id + 3, UUID: "whllo"}

		if err := client.WriteJSON(data); err != nil {
			fmt.Println("error ", err.Error())
		}
		time.Sleep(time.Second * 5)
	}
}

func hello(ctx *gin.Context) {
	type Hello struct {
		ID   uint64 `json:"id,string"`
		UUID uint64 `json:"uuid"`
	}
	data := Hello{ID: math.MaxUint64, UUID: math.MaxUint64}
	ctx.JSON(200, data)
}
