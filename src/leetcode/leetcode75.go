package leetcode

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i0, i2, i := 0, len(nums)-1, 0
	for i <= i2 {
		if nums[i] == 2 {
			tmp := nums[i]
			nums[i] = nums[i2]
			nums[i2] = tmp
			i2--
		} else if nums[i] == 0 {
			if i == i0 {
				i++
				i0++
			} else {
				tmp := nums[i]
				nums[i] = nums[i0]
				nums[i0] = tmp
				i0++
			}
		} else {
			i++
		}
	}
}
