package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	. "outback/leetcode/common"
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
func isRectangleCover(rectangles [][]int) bool {
	exitMap := make(map[string]int)
	area := 0
	for _, rect := range rectangles {
		area += Area(rect)
		fourRes := four(rect)
		fmt.Println("fourres", fourRes)
		for _, s := range fourRes {
			key := fmt.Sprintf("%d_%d", s[0], s[1])
			if _, exist := exitMap[key]; exist {
				exitMap[key]--
			} else {
				exitMap[key]++
			}
			//fmt.Println(exitMap)
		}
	}
	remainder := make([][]int, 0)

	for key, v := range exitMap {
		if v != 0 {
			//fmt.Println("key is ",key)
			remainder = append(remainder, parse(key))
		}
	}
	fmt.Println("remiaider", remainder, area)

	if len(remainder) != 4 {
		return false
	}
	// 计算面积
	a := 0

	for _, rem1 := range remainder {
		for _, rem2 := range remainder {
			a += Area2(rem1, rem2)
		}
	}
	fmt.Println(a)
	if a == area {
		return true
	}
	return false
}

func four(s []int) [][]int {
	res := make([][]int, 4)
	res[0] = []int{s[0], s[1]}
	res[1] = []int{s[0], s[3]}
	res[2] = []int{s[2], s[1]}
	res[3] = []int{s[2], s[3]}
	return res
}
func Area(s []int) int {
	return (s[2] - s[0]) * (s[3] - s[1])
}
func Area2(s [][]int) int {
	x, y := 0, 0

}

func parse(s string) []int {
	split := strings.Split(s, "_")
	res := make([]int, 0)
	for _, sp := range split {
		i, _ := strconv.Atoi(sp)
		res = append(res, i)
	}
	return res
}
