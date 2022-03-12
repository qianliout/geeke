package main

import (
	"fmt"
	"strings"
)

func main() {
	path := simplifyPath("/a/../../b/../c//.//")
	fmt.Println("apth ", path)
}

func simplifyPath(path string) string {
	stark := make([]string, 0)
	split := strings.Split(path, "/")
	for i := range split {
		if split[i] == ".." {
			if len(stark) > 0 {
				stark = stark[:len(stark)-1]
			}
			continue
		}
		if split[i] == "/" || split[i] == "" || split[i] == "." {
			continue
		}
		stark = append(stark, split[i])
	}
	if len(stark) == 0 {
		return "/"
	}

	return "/" + strings.Join(stark, "/")
}
