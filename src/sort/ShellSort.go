package sort

func shellSort(nums []int) {
	step := len(nums) / 2
	for step > 0 {
		for i := 1; i < len(nums); i++ {
			j := i - step
			x := nums[i]
			nums[i] = nums[j]
			for j >= 0 && x < nums[j] {
				nums[j+step] = nums[j]
				j -= step
			}
			nums[j+step] = x
		}
		step--
	}
}
