package starkqueue

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:
输入: "()"
输出: true
示例 2:
输入: "()[]{}"
输出: true
示例 3:
输入: "(]"
输出: false
示例 4:
输入: "([)]"
输出: false
示例 5:
输入: "{[]}"
输出: true
*/
func IsValidParentheses(parentheses string) bool {
	validMap := map[byte]byte{')': '(', '}': '{', ']': '['}

	stark := make([]byte, 0)

	for _, value := range []byte(parentheses) {
		if value == '(' || value == '{' || value == '[' {
			stark = append(stark, value)
		} else {
			second, _ := validMap[value]
			if len(stark) == 0 {
				return false
			}
			if len(stark) > 0 && stark[len(stark)-1] != second {
				return false
			} else {
				stark = stark[:len(stark)-1]
			}
		}
	}
	if len(stark) == 0 {
		return true
	} else {
		return false
	}

}
