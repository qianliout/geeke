package main

import (
	"fmt"
)

func main() {
	primes := countPrimes(5000000)
	primes2 := countPrimes2(5000000)
	fmt.Println("primes is ", primes, primes2)
}

// 给定整数 n ，返回 所有小于非负整数 n 的质数的数量
// 会超时
func countPrimes(n int) int {
	exit := make(map[int]bool)

	for i := 2; i < n; i++ {
		exit[i] = true
	}
	for i := 2; i*i < n; i++ {
		if exit[i] {
			for j := i * i; j < n; j = j + i {
				exit[j] = false
			}
		}
	}
	ans := 0
	for _, v := range exit {
		if v {
			ans++
		}
	}

	return ans
}

func countPrimes2(n int) int {
	primes := make([]bool, n)
	for i := range primes {
		primes[i] = true
	}
	prime := make([]int, 0)
	for i := 2; i < n; i++ {
		if primes[i] {
			prime = append(prime, i)
		}
		for _, p := range prime {
			if i*p >= n {
				break
			}
			primes[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(prime)
}
