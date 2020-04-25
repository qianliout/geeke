package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := compareVersion("1", "1.0.0.0.1")
	fmt.Println("res is ", res)
}

/*
比较两个版本号 version1 和 version2。
如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。
你可以假设版本字符串非空，并且只包含数字和 . 字符。
 . 字符不代表小数点，而是用于分隔数字序列。
例如，2.5 不是“两个半”，也不是“差一半到三”，而是第二版中的第五个小版本。
你可以假设版本号的每一级的默认修订版号为 0。例如，版本号 3.4 的第一级（大版本）
和第二级（小版本）修订号分别为 3 和 4。其第三级和第四级修订号均为 0。
示例 1:
输入: version1 = "0.1", version2 = "1.1"
输出: -1
示例 2:
输入: version1 = "1.0.1", version2 = "1"
输出: 1
示例 3:
输入: version1 = "7.5.2.4", version2 = "7.5.3"
输出: -1
示例 4：
输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，“01” 和 “001” 表示相同的数字 “1”。
示例 5：
输入：version1 = "1.0", version2 = "1.0.0"
输出：0
解释：version1 没有第三级修订号，这意味着它的第三级修订号默认为 “0”。
*/

func compareVersion(version1 string, version2 string) int {
	split1 := strings.Split(version1, ".")
	split2 := strings.Split(version2, ".")
	len1 := len(split1)
	len2 := len(split2)
	if len1 > len2 {
		for i := 0; i <= len1-len2; i++ {
			split2 = append(split2, "0")
		}
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			split1 = append(split1, "0")
		}
	}
	for i := 0; i < len(split1); i++ {
		ato1, _ := strconv.Atoi(split1[i])
		ato2, _ := strconv.Atoi(split2[i])
		if ato1 > ato2 {
			return 1
		} else if ato1 < ato2 {
			return -1
		}
	}
	return 0
}
