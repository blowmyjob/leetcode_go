package sort

func partition(array []int, i int, j int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[i]
	for i < j {
		for j > i && array[j] > pivot {
			j--
		}
		if j > i {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < pivot {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = pivot
	return i
}

func Quicksort(array []int, low int, high int) {
	var pivotPos int //划分基准元素索引
	if low < high {
		pivotPos = partition(array, low, high)
		Quicksort(array, low, pivotPos-1)
		Quicksort(array, pivotPos+1, high)
	}
}