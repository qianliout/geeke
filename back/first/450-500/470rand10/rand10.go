package main

import (
	"fmt"
)

func main() {
	a := 10
	defer func() {
		fmt.Println(a)
	}()
	a++
	return
}

func Print(a int) {
	fmt.Println(a)
}

/*
已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
不要使用系统的 Math.random() 方法。
示例 1:
输入: 1
输出: [7]
示例 2:
输入: 2
输出: [8,4]
示例 3:
输入: 3
输出: [8,1,10]
*/
func rand10() int {
	num := (rand7()-1)*7 + rand7()
	for num > 40 {
		num = (rand7()-1)*7 + rand7()
	}
	//  返回的是1-10
	return 1 + num%10
}

func rand7() int {
	return 0
}
