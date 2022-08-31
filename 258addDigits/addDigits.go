package main

func main() {

}

func addDigits(num int) int {
	for num >= 10 {
		num = resolve(num)
	}
	return num
}

func resolve(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num = num / 10
	}
	return ans
}
