package sort

func bubbleSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j-1] {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
	}
}
