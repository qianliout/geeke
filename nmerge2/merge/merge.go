package merge

import (
	"fmt"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

// 内存排序
func InmSort(in <-chan int) <-chan int {
	out := make(chan int, 0)

	go func() {
		ints := make([]int, 0)
		for n := range in {
			ints = append(ints, n)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))
		// Sort
		sort.Ints(ints)
		fmt.Println("Sort done:", time.Now().Sub(startTime))
		for _, n := range ints {
			out <- n
		}
		close(out)
	}()
	return out
}

// 两路归并排序
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 0)
	go func() {
		n1, ok1 := <-in1
		n2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && n1 <= n2) { // 这里的判断很重要
				out <- n1
				n1, ok1 = <-in1
			} else {
				out <- n2
				n2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

// 多路归并
func MergeN(ints ...<-chan int) <-chan int {
	if len(ints) == 1 {
		return ints[0]
	}
	mid := len(ints) / 2
	return Merge(MergeN(ints[:mid]...), MergeN(ints[mid:]...))
}
