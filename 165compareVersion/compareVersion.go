package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	v1 := "1.01"

	v2 := "1.001"

	version := compareVersion(v1, v2)
	fmt.Println(version)
}

func compareVersion(version1 string, version2 string) int {
	stark1, stark2 := make([]int, 0), make([]int, 0)
	split1 := strings.Split(version1, ".")
	for i := range split1 {
		if split1[i] != "" {
			a, err := strconv.Atoi(split1[i])
			if err == nil {
				stark1 = append(stark1, a)
			}
		}
	}
	split := strings.Split(version2, ".")

	for i := range split {
		if split[i] != "" {
			a, err := strconv.Atoi(split[i])
			if err == nil {
				stark2 = append(stark2, a)
			}
		}
	}

	i := 0

	for i < len(stark1) && i < len(stark2) {
		if stark1[i] < stark2[i] {
			return -1
		} else if stark1[i] > stark2[i] {
			return 1
		} else {
			i++
		}
	}
	if len(stark1) >= i {
		return 1

	}
	return -1
}
