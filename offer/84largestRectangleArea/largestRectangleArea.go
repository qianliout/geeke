package main

func main() {

}

// 单调栈的用法
func largestRectangleArea(heights []int) int {
	ans := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	stark := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		for len(stark) > 0 && heights[stark[len(stark)-1]] > heights[i] {
			height := heights[stark[len(stark)-1]]
			stark = stark[:len(stark)-1]
			right := i - 1
			left := stark[len(stark)-1] + 1
			if ans < (right-left+1)*height {
				ans = (right - left + 1) * height
			}
		}
		stark = append(stark, i)
	}
	return ans
}
