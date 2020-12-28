package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := ResrictResource{
		Id:         3,
		EntId:      34,
		Value:      "hellowrod",
		SourceType: 12,
		Product:    "meiqia",
	}
	bytes, _ := json.Marshal(s)

	fmt.Println(string(bytes))
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
