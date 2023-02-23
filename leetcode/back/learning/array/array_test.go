package array

import (
	"testing"

	"qianliout/leetcode/leetcode/common/array"
)

func TestReverseLinkedList(t *testing.T) {
	root := array.NewLinkedList([]int{2, 3, 4, 5, 7})
	r := ReverseLinkedList(root)
	array.PrintLinkedList(r)
}
