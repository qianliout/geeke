package main

func main() {

}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/
func isValid(s string) bool {
	stark := make([]byte, 0)
	bracket := map[byte]byte{')': '(', ']': '[', '}': '{'}

	ss := []byte(s)
	for _, b := range ss {
		if b == '(' || b == '{' || b == '[' {
			stark = append(stark, b)
		} else {
			if len(stark) == 0 {
				return false
			}
			last := stark[len(stark)-1]
			if bracket[b] != last {
				return false
			}
			stark = stark[:len(stark)-1]
		}
	}
	return len(stark) == 0
}
