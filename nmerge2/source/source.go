package source

import (
	"math/rand"
)

func RandomSource(count int) <-chan int {
	out := make(chan int, 0)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
	}()
	return out
}
