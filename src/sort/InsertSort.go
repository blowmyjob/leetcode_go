package sort

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		x := nums[i]
		nums[i] = nums[j]
		for j >= 0 && x < nums[j] {
			nums[j] = nums[j-1]
			j--
		}
		nums[j+1] = x
	}
}
