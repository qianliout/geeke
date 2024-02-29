package main

import (
	"time"
)

func main() {
	// // var redhatRe = regexp.MustCompile(`(.*) release (\d[\d\.]*)(.*)`)
	// var bclinuxBaseOERe = regexp.MustCompile(`(.*) release (\d[\d\.]*)(.*)`)
	//
	// // str := "BigCloud Enterprise Linux release 7.8.2105 (Core)"
	// // str := "BigCloud Enterprise Linux For Euler release 22.10 LTS"
	// str := "BigCloud Enterprise Linux For Euler release 22.10 (list)"
	// submatch := bclinuxBaseOERe.FindStringSubmatch(str)
	// for i := range submatch {
	// 	fmt.Println(submatch[i])
	// }
	// fmt.Println(len(submatch))
	// reg1 := regexp.MustCompile("\\.p12$")
	// str1 := "/usr/share/terminfo/p/p12"
	// allString := reg1.FindString(str1)
	// fmt.Println(allString)
	//
	// reg2 := regexp.MustCompile(".p12$")
	// allString2 := reg2.FindString(str1)
	// fmt.Println(allString2)
	// str3 := "./usr/share/groff/1.22.4/tmac/hyphenex.us"
	// fmt.Println(strings.Replace(str3, "./", "/", 1))
	//
	// h := strings.HasPrefix("./usr/share/groff/1.22.4/tmac/hyphenex.us", "./")
	// fmt.Println(h)

	time.Sleep(1 * time.Second)
	n := Name{
		ID: "id",
		Na: Hello{Name: "name"},
	}

}

type Hello struct {
	Name string
}

type Name struct {
	ID string
	Hello
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
*/

func twoSum(nums []int, target int) []int {

	exit := make(map[int]int)

	for i, v := range nums {
		if pre, ok := exit[target-v]; ok {
			return []int{i, pre}
		}
		exit[v] = i
	}
	return []int{}
}
