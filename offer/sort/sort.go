package main

func main() {

}

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivotIndex := partitionv2(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

// 双边循环，从右侧开始
func partition(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)

	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	return left
}

// 单边循环
func partitionv2(arr []int, startIndex, endIndex int) int {
	var (
		mark  = startIndex
		pivot = arr[startIndex]
		point = startIndex + 1
	)

	for point < len(arr) {
		if arr[point] < pivot {
			mark++
			arr[mark], arr[point] = arr[point], arr[mark]
		}
		point++
	}

	arr[startIndex], arr[mark] = arr[mark], arr[startIndex]
	return mark
}

func QuickSortNonRecursive(arr []int, startIndex, endIndex int) {
	var (
		s     = v1.NewStack()
		m     = make(map[string]int)
		start = "start_index"
		end   = "end_index"
	)
	m[start] = startIndex
	m[end] = endIndex
	s.Push(m)
	for !s.IsEmpty() {
		param := s.Pop().(map[string]int)
		pivotIndex := partitionv2(arr, param[start], param[end])
		if param[start] < pivotIndex-1 {
			leftParam := make(map[string]int)
			leftParam[start] = param[start]
			leftParam[end] = pivotIndex - 1
			s.Push(leftParam)
		}
		if param[end] > pivotIndex+1 {
			rightParam := make(map[string]int)
			rightParam[start] = pivotIndex + 1
			rightParam[end] = param[end]
			s.Push(rightParam)
		}
	}
}
