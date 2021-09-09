package leetcode

func canJump(nums []int) bool {
	maxJump := 0
	for k, v := range nums {
		if k > maxJump {
			return false
		}
		if k+v > maxJump {
			maxJump = k + v
		}
	}
	return true
}
