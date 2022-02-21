package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hell := Hello{AgeName: "nihaod"}
	bys, _ := json.Marshal(hell)
	fmt.Println(string(bys))
}

type Hello struct {
	AgeName string
}
