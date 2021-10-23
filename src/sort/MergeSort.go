package sort

// 合并
func merge(left []int, right []int) []int {
	leftIdx, rightIdx := 0, 0
	result := []int{}
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] <= right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx += 1
		} else {
			result = append(result, right[rightIdx])
			rightIdx += 1
		}
	}
	// 上述for循环有两种情况：
	// 1.left数组和right数组长度一致
	// 2.left数组和right数组长度不一致，则必定有一方有多余元素，需要把这些多余元素复制到队尾
	result = append(result, left[leftIdx:]...)
	result = append(result, right[rightIdx:]...)
	return result
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	split := n / 2
	left := mergeSort(arr[:split])
	right := mergeSort(arr[split:])
	return merge(left, right)
}
