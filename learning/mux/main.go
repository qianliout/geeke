package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now().Unix()
	userwait()
	fmt.Println(time.Now().Unix() - start)

}

func userwait() {
	num, wg, mux := new(int), new(sync.WaitGroup), new(sync.Mutex)
	wg.Add(2)

	go func(num *int, wg *sync.WaitGroup, mux *sync.Mutex) {
		defer wg.Done()
		for i := 1; i <= 100000000; i++ {
			mux.Lock()
			*num++
			mux.Unlock()
		}
	}(num, wg, mux)

	go func(num *int, wg *sync.WaitGroup, mux *sync.Mutex) {
		defer wg.Done()
		for i := 1; i <= 100000000; i++ {
			mux.Lock()
			*num--
			mux.Unlock()
		}
	}(num, wg, mux)

	wg.Wait()
}

func main2() {
	num, wg, mux := new(int), new(sync.WaitGroup), make(chan bool, 1)
	defer close(mux)
	wg.Add(2)

	go func(num *int, wg *sync.WaitGroup, mux chan bool) {
		defer wg.Done()
		for i := 1; i <= 100000000; i++ {
			mux <- true
			*num++
			<-mux
		}
	}(num, wg, mux)

	go func(num *int, wg *sync.WaitGroup, mux chan bool) {
		defer wg.Done()
		for i := 1; i <= 100000000; i++ {
			mux <- true
			*num--
			<-mux
		}
	}(num, wg, mux)

	wg.Wait()
}
