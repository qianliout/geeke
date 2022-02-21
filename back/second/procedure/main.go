package main

import (
	"fmt"
	"runtime"
	"sort"
)

var all int

func main() {
	r := runFuncName()
	fmt.Println("maian:", r)

	ss := []*Student{{
		Name:   "ab",
		Age:    3,
		Online: false,
	},
		{

			Name:   "bc",
			Age:    4,
			Online: true,
		},
		{

			Name:   "d",
			Age:    2,
			Online: false,
		},
		{

			Name:   "f",
			Age:    1,
			Online: true,
		},
	}
	sort.Sort(Students(ss))
	for i := range ss {
		fmt.Printf("%d,%s,%t\n", ss[i].Age, ss[i].Name, ss[i].Online)
	}
	fmt.Println("all is ", all)
}

type Student struct {
	Name   string
	Age    int
	Online bool
}

type Students []*Student

func (m Students) Len() int {
	return len(m)
}

func (m Students) Less(i, j int) bool {
	all++
	if m[i].Online && !m[j].Online {
		return true
	} else if !m[i].Online && m[j].Online {
		return false
	}
	if m[i].Age > m[j].Age {
		return true
	} else if m[j].Age > m[i].Age {
		return false
	}
	return m[i].Name > m[j].Name
}

func (m Students) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
