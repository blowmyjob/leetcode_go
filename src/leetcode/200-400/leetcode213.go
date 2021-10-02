package leetcode

func rob(nums []int) int {
	var myRob func(nums []int, start, end int) int
	myRob = func(nums []int, start, end int) int {
		pre, cur := 0, 0
		for i := start; i <= end; i++ {
			tmp := cur
			if pre+nums[i] > cur {
				cur = pre + nums[i]
			}
			pre = tmp
		}
		return cur
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 0 {
		return nums[0]
	}
	l1 := myRob(nums, 0, len(nums)-2)
	l2 := myRob(nums, 1, len(nums)-1)
	if l1 > l2 {
		return l1
	}
	return l2
}
