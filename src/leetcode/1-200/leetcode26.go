package leetcode

func removeDuplicates(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return 0
	}
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
