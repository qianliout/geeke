package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "  hello world  "
	ans := reverseWords(s)
	fmt.Println("ans is ", ans)

}

func reverseWords(s string) string {

	stark, ans := make([]string, 0), make([]string, 0)

	split := strings.Split(s, " ")
	for i := range split {
		if split[i] != " " && split[i] != "" {
			stark = append(stark, split[i])
		}
	}

	for i := len(stark) - 1; i >= 0; i-- {
		ans = append(ans, stark[i])
	}
	return strings.Join(ans, " ")
}
