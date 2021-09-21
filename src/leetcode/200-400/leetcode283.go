package leetcode

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i0, i1 := 0, 0
	for i1 < len(nums) {
		if nums[i1] != 0 {
			nums[i0], nums[i1] = nums[i1], nums[i0]
			i0++
		}
		i1++
	}
}
