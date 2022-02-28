package main

func main() {

}

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）。
*/
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var pos float64

	for n > 0 {
		if n&1 == 1 {
			pos = pos * x
		}
		pos = pos * pos
		n = n >> 1
	}
	return pos
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return recursion(x, n)
}

func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := recursion(x, n>>1)
	if n&1 == 1 {
		return res * res * x
	}
	return res * res
}
