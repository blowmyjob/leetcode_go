package leetcode

import "math"

func integerReplacement(n int) int {
	if n == math.MaxInt8 {
		return 32
	}
	if n <= 3 {
		return n - 1
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else {
		return minReplace(integerReplacement(n-1), integerReplacement(n+1)) + 1
	}
}

func minReplace(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
