package contest

func SumOfBeauties(nums []int) int {
	leftNum, rightNum := make([]int, len(nums)), make([]int, len(nums))
	copy(leftNum, nums)
	copy(rightNum, nums)
	for i := 1; i < len(nums); i++ {
		if leftNum[i-1] > leftNum[i] {
			leftNum[i] = leftNum[i-1]
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if rightNum[i+1] < rightNum[i] {
			rightNum[i] = rightNum[i+1]
		}
	}
	res := 0
	for i := 1; i < len(nums)-1; i++ {
		if leftNum[i-1] < nums[i] && nums[i] < rightNum[i+1] {
			res += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res += 1
		}
	}
	return res
}
