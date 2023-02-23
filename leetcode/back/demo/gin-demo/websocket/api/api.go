package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pretty66/websocketproxy"
)

func main() {

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route("/ws", func(r chi.Router) {
				r.Get("/scanner", scannerWs())
			})
		})
	},
	)
	se := &http.Server{
		Addr:    ":8030",
		Handler: r,
	}
	if err := se.ListenAndServe(); err != nil {
		fmt.Println("error is ", err)
	}
}

func scannerWs() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("收到请求")
		newUrl := "ws://127.0.0.1:8081/ws/scanner"
		// 代理目标地址
		// 可设置握手前回调函数，修改request信息
		// ，增减头部，权限验证等等
		wp, err := websocketproxy.NewProxy(newUrl, func(r *http.Request) error {
			fmt.Println("收到转发")
			return nil
		})
		if err != nil {
			return
		}
		wp.ServeHTTP(w, r)
	}
}
