package offer

func Find(target int, array [][]int) bool {
	// write code here
	m, n := 0, len(array[0])-1
	for m < len(array) && n >= 0 {
		if array[m][n] > target {
			n--
		} else if array[m][n] < target {
			m++
		} else {
			return true
		}
	}
	return false
}
