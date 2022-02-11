package segment

type Segment struct {
	Data []int
	Tree []int
}

func NewSegment(nums []int) Segment {
	s := new(Segment)
	s.Data = nums
	s.Tree = buildTree(nums)
	return *s
}

func buildTree(nums []int) []int {
	n := len(nums)
	tree := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		j := i - n
		tree[i] = nums[j]
	}

	for i := n - 1; i > 0; i-- { // 注意的是，这里最开始的元素无意义
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	return tree
}

func (s *Segment) Update(index, val int) {
	s.Data[index] = val
	n := len(s.Data)
	index += n
	s.Tree[index] = val
	for index > 0 {
		left, right := index, index
		if index&1 == 0 {
			right = index + 1
		} else {
			left = index - 1
		}
		// 这里要注意，这里会把tree的第0个数赋值，但是这个值和第1个值是一们的，没有意义
		s.Tree[index/2] = s.Tree[left] + s.Tree[right]
		index = index >> 1
	}
}

// 注意的是，这里的left,right都是包含的
func (s *Segment) SumRange(left, right int) int {
	n := len(s.Data)
	left += n
	right += n
	sum := 0
	for left <= right {
		if left&1 == 1 {
			sum += s.Tree[left]
			left++
		}
		if right&1 == 0 {
			sum += s.Tree[right]
			right--
		}
		left = left >> 1
		right = right >> 1
	}
	return sum
}
