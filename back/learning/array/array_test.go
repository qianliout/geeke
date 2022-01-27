package array

import (
	"testing"

	array2 "qianliout/leetcode/back/common/array"
)

func TestReverseLinkedList(t *testing.T) {
	root := array2.NewLinkedList([]int{2, 3, 4, 5, 7})
	r := ReverseLinkedList(root)
	array2.PrintLinkedList(r)
}
