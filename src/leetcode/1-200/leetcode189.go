package leetcode

func rotate(nums []int, k int) {
	var gcd func(start, end int)
	gcd = func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end++
		}
	}
	n := len(nums)
	k = k % n
	gcd(0, n-1)
	gcd(0, k-1)
	gcd(k, n-1)
}
