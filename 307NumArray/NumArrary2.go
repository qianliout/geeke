package main

type NumArray struct {
	Root *Tree
}

func Constructor(nums []int) NumArray {
	return NumArray{Root: buildTree(nums, 0, len(nums)-1)}
}

func (this *NumArray) Update(index int, val int) {
	updateTree(this.Root, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return query(this.Root, left, right)
}

type Tree struct {
	start int
	end   int
	sum   int
	left  *Tree
	right *Tree
}

func buildTree(nums []int, start, end int) *Tree {
	if start > end {
		return nil
	}
	if start == end {
		return &Tree{
			start: start,
			end:   end,
			sum:   nums[start],
		}
	}
	mid := start + (end-start)/2

	left := buildTree(nums, start, mid)
	right := buildTree(nums, mid+1, end)

	return &Tree{
		start: start,
		end:   end,
		sum:   left.sum + right.sum,
		left:  left,
		right: right,
	}
}

func query(root *Tree, left, right int) int {
	if root.start == left && root.end == right {
		return root.sum
	}
	mid := root.start + (root.end-root.start)/2
	if left > mid {
		return query(root.right, left, right)
	} else if right <= mid {
		return query(root.left, left, right)
	} else {
		return query(root.left, left, mid) + query(root.right, mid+1, right)
	}
}

func updateTree(root *Tree, index, val int) {
	if root.start == root.end {
		root.sum = val
		return
	}
	mid := root.start + (root.end-root.start)/2
	if index <= mid {
		updateTree(root.left, index, val)
	} else {
		updateTree(root.right, index, val)
	}
	root.sum = root.left.sum + root.right.sum
}
