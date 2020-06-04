package main

func main() {

}

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
示例：
s = "leetcode"
返回 0
s = "loveleetcode"
返回 2
提示：你可以假定该字符串只包含小写字母。
*/

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	uMap := make(map[byte]int)

	for _, v := range []byte(s) {
		uMap[v]++
	}

	for i, v := range []byte(s) {
		if uMap[v] <= 1 {
			return i
		}
	}
	return -1
}
