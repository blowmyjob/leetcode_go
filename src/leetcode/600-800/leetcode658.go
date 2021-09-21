package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - 1

	removeNums := len(arr) - k
	for removeNums != 0 {
		if x-arr[start] <= arr[end]-x {
			end--
		} else {
			start++
		}

		removeNums--
	}

	return arr[start : end+1]
}
