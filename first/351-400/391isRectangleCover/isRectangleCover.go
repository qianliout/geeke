package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	tri := [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}
	cover := isRectangleCover(tri)
	fmt.Println("res is ", cover)
}

var regPhone *regexp.Regexp = regexp.MustCompile("^1[0-9]{10}")

func ValidPhone(phone string) bool {
	phone = strings.Trim(phone, " ")
	fmt.Println(phone)
	return regPhone.MatchString(phone)
}

/*
我们有 N 个与坐标轴对齐的矩形, 其中 N > 0, 判断它们是否能精确地覆盖一个矩形区域。
每个矩形用左下角的点和右上角的点的坐标来表示。例如， 一个单位正方形可以表示为 [1,1,2,2]。 ( 左下角的点的坐标为 (1, 1)
以及右上角的点的坐标为 (2, 2) )。
示例 1:
rectangles = [
  [1,1,3,3],
  [3,1,4,2],
  [3,2,4,4],
  [1,3,2,4],
  [2,3,3,4]
]
返回 true。5个矩形一起可以精确地覆盖一个矩形区域。
示例 2:
rectangles = [
  [1,1,2,3],
  [1,3,2,4],
  [3,1,4,2],
  [3,2,4,4]
]
返回 false。两个矩形之间有间隔，无法覆盖成一个矩形。
示例 3:
rectangles = [
  [1,1,3,3],
  [3,1,4,2],
  [1,3,2,4],
  [3,2,4,4]
]
返回 false。图形顶端留有间隔，无法覆盖成一个矩形。
示例 4:
rectangles = [
  [1,1,3,3],
  [3,1,4,2],
  [1,3,2,4],
  [2,2,4,4]
]
返回 false。因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。
*/

// 思路很简单，满足三个条件就可以
// 1,最外面的点只能是2个
// 2,里面的点要么是2个，要么是4个
// 最外的点所围成的面积和所有的面积之各一致
// 这道题，题目给的条件就是我们平常学的直角坐标系
func isRectangleCover(rectangles [][]int) bool {
	up, down, left, right := 0, math.MaxInt64, math.MaxInt64, 0
	area := 0
	pointMap := make(map[[2]int]int)

	for _, rectangle := range rectangles {
		up = Max(up, rectangle[3])
		down = Min(down, rectangle[1])
		left = Min(left, rectangle[0])
		right = Max(right, rectangle[2])
		area += (rectangle[3] - rectangle[1]) * (rectangle[2] - rectangle[0])
		// 这是原来的两个点（左下，和右上）
		pointMap[[2]int{rectangle[0], rectangle[1]}]++
		pointMap[[2]int{rectangle[2], rectangle[3]}]++

		// 补下剩下的两个点，左上和右下
		pointMap[[2]int{rectangle[0], rectangle[3]}]++
		pointMap[[2]int{rectangle[2], rectangle[1]}]++

	}
	// 下面进行判断了
	// fmt.Println("rare is ", area, (up-down)*(right-left))
	if (up-down)*(right-left) != area {
		return false
	}
	// 最顶上的四个点只会出现一次，这里加一次，刚好符合下面for循环中的判断
	pointMap[[2]int{left, down}]++
	pointMap[[2]int{left, up}]++
	pointMap[[2]int{right, down}]++
	pointMap[[2]int{right, up}]++

	for _, v := range pointMap {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}
