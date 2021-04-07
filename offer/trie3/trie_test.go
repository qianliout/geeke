package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkTrie(b *testing.B) {
	tr := Constructor()
	all := b.N
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := gen(rd, all)
	b.ResetTimer()

	for n := 0; n < all; n++ {
		tr.Insert(res[n])
	}

	for n := 0; n < all; n++ {
		tr.Delete(res[n])
	}

	for n := 0; n < all; n++ {
		tr.Search(res[n])
	}
}

func gen(rd *rand.Rand, n int) [][]uint8 {
	res := make([][]uint8, n)
	for i := 0; i < n; i++ {
		l := rd.Intn(32)
		ans := make([]uint8, 0)
		for j := 0; j < l; j++ {
			ans = append(ans, uint8(rd.Intn(2)))
		}
		res[i] = ans
	}
	return res
}
