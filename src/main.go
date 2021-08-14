package main

import (
	"leetcode/src/contest"
)

func main() {
	num := [][]byte{
		{'.', '.', 'W', '.', 'B', 'W', 'W', 'B'},
		{'B', 'W', '.', 'W', '.', 'W', 'B', 'B'},
		{'.', 'W', 'B', 'W', 'W', '.', 'W', 'W'},
		{'W', 'W', '.', 'W', '.', '.', 'B', 'B'},
		{'B', 'W', 'B', 'B', 'W', 'W', 'B', '.'},
		{'W', '.', 'W', '.', '.', 'B', 'W', 'W'},
		{'B', '.', 'B', 'B', '.', '.', 'B', 'B'},
		{'.', 'W', '.', 'W', '.', 'W', '.', 'W'},
	}
	contest.CheckMove(num, 5, 4, 'W')
}
