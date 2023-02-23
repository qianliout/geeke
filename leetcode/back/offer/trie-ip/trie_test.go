package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func BenchmarkTrie(b *testing.B) {
	tr := Constructor()
	all := b.N
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := genIP(rd, all)
	fmt.Println("res length", len(res), all)
	b.ResetTimer()

	// 这里验证正确性也很有必要，验证过程如下
	// 1,先插入数据，再去查看数据，如果有没有查到的就说明程序不正确
	// 2,先插入数据，再删除部分数据，再去查找已删除数据，如果查到了就说明程序不正确
	for n := 0; n < len(res); n++ {
		tr.Insert(res[n])
	}

	for n := 0; n < len(res); n += 2 {
		tr.Delete(res[n])
	}

	for n := 0; n < len(res); n += 2 {
		if query, err := tr.Query(res[n]); err == nil && query {
			fmt.Println(query, err)
		}
	}
}

// 模拟并发
func BenchmarkTrieConcurrent(b *testing.B) {
	tr := Constructor()
	all := b.N
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := genIP(rd, all)
	// fmt.Println("res length", res)
	b.ResetTimer()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			insert_test(&tr, &wg, res)
		}()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			update_test(&tr, &wg, res)
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			query_test(&tr, &wg, res)
		}()
	}
	wg.Wait()
}

func insert_test(tr *Trie, wg *sync.WaitGroup, ips []string) {
	for n := 0; n < len(ips); n++ {
		tr.Insert(ips[n])
	}
	wg.Done()
}

func update_test(tr *Trie, wg *sync.WaitGroup, ips []string) {
	for n := 0; n < len(ips); n++ {
		tr.Update(ips)
	}
	wg.Done()
}

func query_test(tr *Trie, wg *sync.WaitGroup, ips []string) {
	for n := 0; n < len(ips); n++ {
		// 这里注要做性能测试，所以不关心执行结果
		tr.Query(ips[n])
	}
	wg.Done()
}

func TestUint8ToByte(t *testing.T) {
	res := uint8ToUint8slice(34)
	fmt.Println(res)
}

func TestIpToUint8Slice(t *testing.T) {
	res := ipToUint8Slice("255.2.0.0/10")
	fmt.Println(res)
}

func genIP(rd *rand.Rand, n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		ips := make([]string, 0)
		for j := 0; j < 4; j++ {
			l := rd.Intn(254) + 1
			ips = append(ips, strconv.Itoa(l))
		}
		ip := strings.Join(ips, ".")
		l := rd.Intn(31) + 1
		ip = ip + "/" + strconv.FormatUint(uint64(l), 10)
		res[i] = ip
	}
	return res
}
