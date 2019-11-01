package dfsbfs

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	Gen(n, 0, 0, "", &result)
	return result
}

func Gen(n, left, right int, res string, result *[]string) {
	// left,right表示已经使用的括号数量,
	// result 表示当前已经生成的结果，
	if left == right && right == n {
		*result = append(*result, res)
	}
	if left < n {
		Gen(n, left+1, right, res+"(", result)
	}
	if right < left && right < n {
		Gen(n, left, right+1, res+")", result)
	}
}

/*
方法三：闭合数
思路
为了枚举某些内容，我们通常希望将其表示为更容易计算的不相交子集的总和。
考虑有效括号序列 S 的 闭包数：至少存在 index >= 0，使得 S[0], S[1], ..., S[2*index+1]是有效的。 显然，每个括号序列都有一个唯一的闭包号。 我们可以尝试单独列举它们。
算法
对于每个闭合数 c，我们知道起始和结束括号必定位于索引 0 和 2*c + 1。然后两者间的 2*c 个元素一定是有效序列，其余元素一定是有效序列。
*/
func GenerateParenthesisByC(n int) []string {
	if n == 0 {
		return []string{""} //这里一定要加一下空字符串才行
	}
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		for _, left := range GenerateParenthesisByC(i) {
			for _, right := range GenerateParenthesisByC(n - i - 1) {
				res := "(" + left + ")" + right
				result = append(result, res)
			}
		}
	}
	return result
}

func generateParenthesisByDP(n int) []string {

}
