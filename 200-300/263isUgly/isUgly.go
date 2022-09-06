package main

func main() {

}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	ugly := []int{2, 3, 5}
	for n != 1 {
		flag := false
		for i := range ugly {
			if n%ugly[i] == 0 {
				n = n / ugly[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return n == 1
}
func isUgly2(n int) bool {
	if n <= 0 {
		return false
	}
	ugly := []int{2, 3, 5}
	for i := range ugly {
		for n%ugly[i] == 0 {
			n = n / ugly[i]
		}
	}
	return n == 1
}
