package main

import (
	"fmt"
)

func main() {

	s := &Student{}
	s.AddAge()
	fmt.Println("age ", s.Age, s.Fim)

	fmt.Println(((1<<2)>>2)&1 == 1)

}

type Student struct {
	Name string
	Age  int
	Fim  []string
}

func (s *Student) AddAge() {
	s.Age++
	s.Fim = append(s.Fim, "heloo")
}
