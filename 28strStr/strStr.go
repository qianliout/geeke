package main

func main() {

}

/*
实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
*/
func strStr(haystack string, needle string) int {
	hh := []byte(haystack)

	for i := 0; i < len(haystack)-len(needle); i++ {
		if string(hh[i:len(needle)]) == needle {
			return i
		}
	}
	return -1
}
