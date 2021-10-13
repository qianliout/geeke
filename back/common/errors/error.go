package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
	}()

	go func() {
		for {
			i, open := <-ch
			fmt.Println(i, open)
		}

	}()
	time.Sleep(time.Second)

}
