package main

import (
	"fmt"
	"time"
)

func main() {
	tokenChan, downChan := make(chan string), make(chan string)
	defer close(tokenChan)
	defer close(downChan)

	go GetToken(tokenChan)
	go GenToken(downChan, tokenChan)

	downChan <- "helloword"
	time.Sleep(time.Second * 5)

}

func GetToken(tokenChan chan string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		token := <-tokenChan
		fmt.Println("get token is ", token)
	}
}

func GenToken(downChan, tokenChan chan string) {
	for {
		select {
		case msg := <-downChan:
			fmt.Println("get downChan is ", msg)
			return
		case tokenChan <- fmt.Sprintf("token-is-%d", time.Now().Unix()):
		}
	}
}
