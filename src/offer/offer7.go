package offer

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}
	first, last := 0, len(rotateArray)-1
	if rotateArray[first] < rotateArray[last] {
		return rotateArray[first]
	}
	for first < last {
		mid := first + (last-first)>>1
		if rotateArray[mid] > rotateArray[last] {
			first = mid + 1
		} else if rotateArray[mid] < rotateArray[last] {
			last = mid
		} else {
			last -= 1
		}
	}
	return rotateArray[first]
}

func jumpFloorII(number int) int {
	// write code here
	if number == 1 {
		return 1
	}
	dp := make([]int, number+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= number; i++ {
		dp[i] = 2 * dp[i-1]
	}
	return dp[number]
}

func rectCover(number int) int {
	// write code here
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	dp := make([]int, number+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
