package main

import (
	"fmt"
	"leetcode/src/leetcode"
)

func main() {
	num := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(leetcode.Compress(num))
}
