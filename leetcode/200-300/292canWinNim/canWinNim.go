package main

func main() {
	n := 1199886170
	canWin2(n)

}

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}
	// 博弈论的思想，如果能保证走不到n-1,n-2,n-3,那我就赢了
	// canWinNic(n-1)表示能走到n-1
	return !canWinNim(n-1) || !canWinNim(n-2) || !canWinNim(n-3)
}

func canWin(n int, exit map[int]bool) bool {
	if n <= 3 {
		return true
	}
	if c, ok := exit[n]; ok {
		return c
	}
	exit[n] = !canWin(n-1, exit) || !canWin(n-2, exit) || !canWin(n-3, exit)
	return exit[n]
}

func canWin2(n int) bool {
	if n <= 3 {
		return true
	}
	a, b, c, can := true, true, true, false
	for i := 4; i <= n; i++ {
		can = !a || !b || !c
		a, b, c = b, c, can

	}
	return can
}

func canWinNim2(n int) bool {
	return n%4 != 0
}
