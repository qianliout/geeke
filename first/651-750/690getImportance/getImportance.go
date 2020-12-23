package main

func main() {

}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// 本质就是一个树的搜索
func getImportance(employees []*Employee, id int) int {
	var root *Employee
	for _, e := range employees {
		if e.Id == id {
			root = e
			break
		}
	}
	// 测试用例中不会有这样的情况，但是判断一下也没有什么错了
	if root == nil {
		return 0
	}
	ans := root.Importance
	for _, e := range root.Subordinates {
		ans += getImportance(employees, e)
	}
	return ans
}
