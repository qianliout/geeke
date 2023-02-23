package main

func main() {

}

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/
func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans = ans ^ n
	}
	return ans
}

/*
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
*/
func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}
func singleNumber2(nums []int) []int {
	diff := 0
	for _, n := range nums {
		diff = diff ^ n
	}
	divide := diff ^ (-diff)
	x := 0
	for _, n := range nums {
		if n&divide == 1 {
			x = x ^ n
		}
	}
	return []int{x, x ^ diff}
}
