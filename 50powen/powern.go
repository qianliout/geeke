package main

import "fmt"

func main() {
	var f float64 = 0.00001

	var n int = 2147483647
	res := myPow(f, n)
	fmt.Println(res)
}

// 递归实现
func myPow(x float64, n int) float64 {
	if n < 0 {
		return float64(1) / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	//return myPow(x, int(n/2)) * myPow(x, n-int(n/2)) // 这种做法会超出时间限制
	if n%2 == 1 {
		return x * myPow(x, n-1)
	} else {
		return myPow(x*x, n/2)
	}
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var pow float64 = 1
	if n == 0 {
		return 1
	}
	for n > 0 {
		if n&1 == 1 { // 说明n是奇数
			pow = pow * x
		}
		x = x * x
		n = n >> 1 // 相当于除以2
	}
	return pow
}
