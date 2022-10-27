package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// TestWebSocket
// @Description: 测试WebSocket脚本
// @param t
func BenchmarkWebSocket(t *testing.B) {
	for n := 0; n < 1000; n++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			url := "ws://console-test-cn.tensorsecurity.cn/api/v2/containerSec/scanner/ws"
			c, res, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				log.Fatal("连接失败:", err)
			}
			log.Printf("响应:%s", fmt.Sprint(res))

			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Fatal(err)
					break
				}
				log.Printf("收到消息: %s", message)
			}
		}(n)

	}
	time.Sleep(time.Minute)
}
