package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
}

func hello() {
	s := []int{1, 2, 3, 4, 5, 6}

	go func() {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == 3 || i == 1 {
				fmt.Println("ss", s[i])
				s = append(s[:i], s[i+1:]...)
			}
			s = append(s, s[i]+10)
		}
	}()
	go func() {
		s = append(s, 20)
	}()
	time.Sleep(time.Second)
	fmt.Println(s)
}
