package merge

import (
	"sort"
)

// 内存排序
func InmSort(in <-chan int64) <-chan int64 {

	out := make(chan int64, 0)
	go func() {
		ints := make([]int, 0)
		defer close(out)
		for n := range in {
			ints = append(ints, int(n))
		}
		sort.Ints(ints)

		for _, n := range ints {
			out <- int64(n)
		}
	}()

	return out
}

// 双路归并
func MergeTwo(ch1, ch2 <-chan int64) <-chan int64 {
	out := make(chan int64, 0)

	go func() {
		defer close(out)
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		for ok1 || ok2 {
			if !ok2 || (ok1 && n1 <= n2) {
				out <- n1
				n1, ok1 = <-ch1
			} else {
				out <- n2
				n2, ok2 = <-ch2
			}
		}
	}()
	return out
}

// 多路归并
func MergeN(chs ...<-chan int64) <-chan int64 {
	if len(chs) == 1 {
		return chs[0]
	}
	mid := len(chs) / 2
	return MergeTwo(MergeN(chs[:mid]...), MergeN(chs[mid:]...))
}
