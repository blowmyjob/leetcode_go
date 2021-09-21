package leetcode

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	maxNum := 0
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}
	if sum%k != 0 || maxNum > sum/k {
		return false
	}
	dp := make([]bool, len(nums))
	return dfsCan(nums, k, 0, sum/k, 0, dp)
}
func dfsCan(num []int, k, start, target, cur int, dp []bool) bool {
	if k == 0 {
		return true
	}
	if target == cur {
		return dfsCan(num, k-1, 0, target, 0, dp)
	}
	for i := start; i < len(num); i++ {
		if !dp[i] && cur+num[i] <= target {
			dp[i] = true
			if dfsCan(num, k, i+1, target, cur+num[i], dp) {
				return true
			}
			dp[i] = false
		}
	}
	return false
}
