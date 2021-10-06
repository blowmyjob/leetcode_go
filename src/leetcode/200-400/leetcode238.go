package leetcode

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	k := 1
	for i := 0; i < len(nums); i++ {
		res[i] = k
		k *= nums[i]
	}
	k = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= k
		k *= nums[i]
	}
	return res
}
