package main

import (
	"fmt"
	"leetcode/src/leetcode"
)

func main() {
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	fmt.Println(leetcode.FindMedianSortedArrays(num1, num2))
}
