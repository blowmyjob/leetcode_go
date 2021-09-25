package sort

func QuickSort(nums []int, left, right int) {
	if left < right {
		l, r := left, right
		i := left
		x := nums[i]
		for l < r {
			for l < r && nums[r] > x {
				r--
			}
			if l < r {
				nums[l] = x
				l++
			}
			for l < r && nums[l] < x {
				l++
			}
			if l < r {
				nums[r] = x
				r--
			}
		}
		nums[i] = x
		QuickSort(nums, l, i-1)
		QuickSort(nums, i+1, r)
	}
}
