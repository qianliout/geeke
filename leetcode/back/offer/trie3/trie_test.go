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

// func BenchmarkTrieConcurrent(b *testing.B) {
// 	tr := Constructor()
// 	all := b.N
// 	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	res := gen(rd, all)
// 	b.ResetTimer()
//
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			insert_test(&tr, &wg, res)
// 		}()
// 	}
//
// 	for i := 0; i < 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			delete_test(&tr, &wg, res)
// 		}()
// 	}
//
// 	for i := 0; i < 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			search_test(&tr, &wg, res)
// 		}()
// 	}
// 	wg.Wait()
// }
//
// func insert_test(tr *Trie, wg *sync.WaitGroup, ips [][]uint8) {
// 	for n := 0; n < len(ips); n++ {
// 		tr.Insert(ips[n])
// 	}
// 	wg.Done()
// }
//
// func delete_test(tr *Trie, wg *sync.WaitGroup, ips [][]uint8) {
// 	for n := 0; n < len(ips); n++ {
// 		tr.Delete(ips[n])
// 	}
// 	wg.Done()
// }
//
// func search_test(tr *Trie, wg *sync.WaitGroup, ips [][]uint8) {
// 	for n := 0; n < len(ips); n++ {
// 		tr.Query(ips[n])
// 	}
// 	wg.Done()
// }
