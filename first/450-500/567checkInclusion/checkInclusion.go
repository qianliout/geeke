package main

func main() {

}

/*
 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 示例1:
 输入: s1 = "ab" s2 = "eidbaooo"
 输出: True
 解释: s2 包含 s1 的排列之一 ("ba").
 示例2:
 输入: s1= "ab" s2 = "eidboaoo"
 输出: False
 注意：
     输入的字符串只包含小写字母
     两个字符串的长度都在 [1, 10,000] 之间
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	preMap := make(map[byte]int)

	for _, b := range []byte(s1) {
		preMap[b]++
	}
	left, right := 0, 0
	ss2 := []byte(s2)
	k := len(s1)

	curMap := make(map[byte]int)
	for right < len(s2) {
		curMap[ss2[right]]++
		if right-left < k-1 {
			right++
		} else {
			equal := areEqual(preMap, curMap)
			if equal {
				return true
			}
			right++
			v := curMap[ss2[left]]
			if v == 1 {
				delete(curMap, ss2[left])
			} else {
				curMap[ss2[left]]--
			}
			left++
		}
	}
	return false
}

func areEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}
