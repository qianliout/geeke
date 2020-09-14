package main

func main() {

}

/*
给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。
注意:
    输入只包含小写英文字母。
    输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。
    输入字符串的长度小于 50,000。
示例 1:
输入: "owoztneoer"
输出: "012" (zeroonetwo)
示例 2:
输入: "fviefuro"
输出: "45" (fourfive)
*/

// 保证输入是合法有
/*
1:"one"

2:"two"

3:"three"

4:"four"

5："five"

6："six"

7:"seven"

8:"eight"

9:"nine"

*/

func originalDigits(s string) string {
	numMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

}
