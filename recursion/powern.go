package recursion

import (
	"math"
)

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
示例 1:
输入: 2.00000, 10
输出: 1024.00000
*/

func MyPow(x float64, n int) float64 {
	return MyPowByRec(x, n)
}

func MyPowByTBD(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func MyPowByFor(x float64, n int) float64 {
	var result float64 = 0
	for i := 0; i < n; i++ {
		result = result * x
	}
	return result
}

func MyPowByRec(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPowByRec(x, -n)
	}
	if n == 1 {
		return x
	}
	//这样写为什么不对，说说原因呢
	//return MyPowByRec(x, int(n/2)) * MyPowByRec(x, n-int(n/2))
	if n%2 == 1 {
		return x * MyPowByRec(x, n-1)
	} else {
		return MyPowByRec(x*x, n/2)
	}
}

//这个代码是很厉害的代码
func MyPowByWhile(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow float64 = 1
	for n > 0 {
		if n&1 == 1 {
			pow = pow * x
		}
		x = x * x
		n = n >> 1
	}
	return pow
}
