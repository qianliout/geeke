package main

import (
	"strings"
)

func main() {

}

type MapSum struct {
	value map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {

	return MapSum{value: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.value[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	ans := 0
	for k, v := range this.value {
		if strings.HasPrefix(k, prefix) {
			ans += v
		}
	}
	return ans
}
