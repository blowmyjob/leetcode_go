package main

import (
	"fmt"
	"leetcode/src/sort"
)

func main() {
	var arr = []int{5, 2, 1, 4, 3}
	sort.Quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
