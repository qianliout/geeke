package binarysearch

/*
这个程序有两个方法要特别注意，一是取mi在值mid := (left + right + 1) >> 1
二是，取right的值
*/
func Sqrt(x int) int {
	left := 0
	right := x
	for left < right {
		mid := (left + right + 1) >> 1
		sqrt := mid * mid
		if sqrt > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
