package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	data2 "qianliout/leetcode/leetcode/back/second/pprof/data"
)

func main() {
	go func() {
		for {
			log.Println(data2.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
