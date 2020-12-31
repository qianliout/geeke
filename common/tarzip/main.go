package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{2, 3}
	b := make([]int, 2)
	copy(b, a)
	b[1] = 5
	fmt.Println(a, b)
}

// 限制资源
type ResrictResource struct {
	Id         int64  `json:"id"`
	EntId      int64  `json:"ent_id"`
	Value      string `json:"value"`
	SourceType uint8  `json:"source_type"`
	Product    string `json:"product"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
