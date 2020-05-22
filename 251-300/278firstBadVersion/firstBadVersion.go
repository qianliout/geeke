package main

import ()

func main() {

}

func firstBadVersion(n int) int {
	if n <= 0 {
		return 0
	}
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		bad := isBadVersion(mid)
		if !bad {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func isBadVersion(version int) bool {}
