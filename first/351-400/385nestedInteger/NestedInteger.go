package main

import (
	"strconv"
	"strings"
)

func main() {

}

type NestedInteger struct {
	integer int
	nexted  *NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (nest NestedInteger) IsInteger() bool {
	if nest.nexted == nil {
		return true
	}
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (nest NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (nest NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (nest NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (nest NestedInteger) GetList() []*NestedInteger {}

/*
字符串只包含数字0-9、[、-、,、]
*/
// 使用递归试试呢

func deserialize(s string) *NestedInteger {

	ss := []byte(s)
	for len(ss) > 0 && ss[0] != '[' {
		ss = ss[1 : len(ss)-1]
	}
	return dfs(ss)
}

func dfs(ss []byte) *NestedInteger {
	nested := new(NestedInteger)
	if ss[0] == '[' {
		ss = ss[1 : len(ss)-1]
		nested.nexted = dfs(ss)
		return nested
	}
	temNum := make([]byte, 0)
	for ss[0] >= '0' && ss[0] <= 9 {
		temNum = append(temNum, ss[0])
		ss = ss[1:]
	}
	n, _ := strconv.Atoi(string(temNum))
	nested.integer = n
	temNum = make([]byte, 0)
	return nested
}
