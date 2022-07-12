package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/rookie-ninja/rk-boot"
)

func main() {
	str := "hello:"
	split := strings.Split(str, ":")
	fmt.Println(len(split), split)

}

type Hello struct {
	mutex sync.Mutex
}

func (h *Hello) Hello1() {
	fmt.Println("Hello1 start is ", time.Now().Unix())
	h.mutex.Lock()
	defer h.mutex.Unlock()
	time.Sleep(time.Second * 4)
	fmt.Println("Hello1 end is ", time.Now().Unix())
}

func (h *Hello) Hello2() {
	fmt.Println("Hello2 start is ", time.Now().Unix())
	h.mutex.Lock()
	defer h.mutex.Unlock()
	time.Sleep(time.Second * 4)
	fmt.Println("Hello2 end is ", time.Now().Unix())
}
