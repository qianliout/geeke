package array

import "testing"

func TestNewLinkedList(t *testing.T) {
	root := NewLinkedList([]int{2, 3, 4, 1, 2})
	PrintLinkedList(root)
}
