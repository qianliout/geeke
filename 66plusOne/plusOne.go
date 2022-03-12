package main

func main() {

}

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func plusOne(digits []int) []int {
	last := 0
	ans := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + last
		if i == len(digits)-1 {
			num += 1
		}

		if num >= 10 {
			last = 1
			num = num - 10
		} else {
			last = 0
		}
		ans[i+1] = num
	}
	if last > 0 {
		ans[0] = last
		return ans
	}
	return ans[1:]
}
