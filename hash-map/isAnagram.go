package hashmap

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
示例 1:
输入: s = "anagram", t = "nagaram"
输出: true
示例 2:
输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。
进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

func IsAnagram(s string, t string) bool {
	return IsAnagramByMap(s, t)
}

func IsAnagramByMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	preMap := make(map[byte]int, 0)
	afterMap := make(map[byte]int, 0)

	for _, value := range []byte(s) {
		if v, ok := preMap[value]; !ok {
			preMap[value] = 1
		} else {
			preMap[value] = v + 1
		}
	}
	for _, value := range []byte(t) {
		if v, ok := afterMap[value]; !ok {
			afterMap[value] = 1
		} else {
			afterMap[value] = v + 1
		}
	}
	for _, value := range []byte(s) {
		preN := preMap[value]
		afterN := afterMap[value]
		if preN != afterN {
			return false
		}
	}
	return true
}

func IsAnagramBySort(s, t string) bool {
	return true

}
