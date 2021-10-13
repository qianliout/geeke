package main

func main() {

}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool           { return true }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	Queue []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	queue := make([]int, 0)
	dfs(nestedList, &queue)
	return &NestedIterator{Queue: queue}

}

func (this *NestedIterator) Next() int {
	n := this.Queue[0]
	this.Queue = this.Queue[1:len(this.Queue)]
	return n

}

func (this *NestedIterator) HasNext() bool {

	return len(this.Queue) > 0
}

func dfs(list []*NestedInteger, res *[]int) {
	for _, node := range list {
		if node.IsInteger() {
			*res = append(*res, node.GetInteger())
		} else {
			dfs(node.GetList(), res)
		}
	}
}
