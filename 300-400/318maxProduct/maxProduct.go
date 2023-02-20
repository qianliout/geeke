package main

func main() {

}

// 会超时
func maxProduct1(words []string) int {
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i - 1; j >= 0; j-- {
			if !diff(words[i], words[j]) {
				if len(words[i])*len(words[j]) > ans {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return ans
}

func mixin(word1, word2 string) bool {
	exit1 := make(map[byte]bool)
	for i := range word1 {
		exit1[word1[i]] = true
	}
	for i := range word2 {
		if exit1[word2[i]] {
			return true
		}
	}
	return false
}

// 使用位移操作
func diff(word1, word2 string) bool {
	b1, b2 := 0, 0
	for _, w := range word1 {
		b1 = b1 | (1 << (w - 'a'))
	}
	for _, w := range word2 {
		b2 = b2 | (1 << (w - 'a'))
	}
	return !(b1&b2 == 0)
}

func maxProduct(words []string) int {
	ans, bits := 0, make([]int, len(words))
	for i := range words {
		b := 0
		for j := range words[i] {
			b = b | (1 << (words[i][j] - 'a'))
		}
		bits[i] = b
	}

	for i := 0; i < len(bits); i++ {
		for j := i - 1; j >= 0; j-- {
			if bits[i]&bits[j] == 0 {
				if ans < len(words[i])*len(words[j]) {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return ans
}
