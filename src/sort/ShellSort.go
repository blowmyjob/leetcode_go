package sort

func shellSort(nums []int) {
	step := len(nums) / 2
	for step > 0 {
		for i := 1; i < len(nums); i++ {
			j := i - step
			x := nums[i]
			nums[i] = nums[j]
			for j >= 0 && x < nums[j] {
				nums[j] = nums[j-1]
				j -= step
			}
			nums[j+1] = x
		}
	}
}
