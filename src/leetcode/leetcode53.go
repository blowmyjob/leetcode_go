package leetcode

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= 0 {
			nums[i] += nums[i-1]
		}
		res = max(nums[i], res)
	}
	return res
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}

func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j, base := left, right, nums[left]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
