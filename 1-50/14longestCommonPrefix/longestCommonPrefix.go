package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println("res is ", res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	end := 0

	slices := make([][]byte, 0)
	minlenght := math.MaxInt64

	for _, i2 := range strs {
		if len(i2) < minlenght {
			minlenght = len(i2)
		}
		slices = append(slices, []byte(i2))
	}

	stop := false
	for end < minlenght && !stop {
		for i := 0; i < len(slices)-1; i++ {
			if slices[i][end] != slices[i+1][end] {
				stop = true
				break
			}
		}
		if stop {
			break
		}
		end++
	}
	return string(slices[0][:end])
}
