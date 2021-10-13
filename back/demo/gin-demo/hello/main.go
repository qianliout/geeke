package main

import (
	"fmt"

	_ "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	t := &Tdu{}
	if t.Age == nil {
		fmt.Println("t.Age is nil")
	}

	fmt.Println(t.Age.GetAge2())
	fmt.Println(t.Age.GetAge())

	// fmt.Println(t.Age.GetAge2())
	// fmt.Println(t.Age.GetAge())
}

type Tdu struct {
	Age *Age
}

type Age struct {
	Hello int
}

func (s *Age) GetAge() int {
	age := s.Hello
	return age
}

func (s *Age) GetAge2() int {

	return 12
}
