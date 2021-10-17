package contest

import (
	"math"
	"sort"
)

func winnerOfGame(colors string) bool {
	if len(colors) <= 2 {
		return false
	}
	a, b := 0, 0
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A' && colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			a++
		} else if colors[i] == 'B' && colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			b++
		}
	}
	return a > b
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	sum := 0
	for i := 0; i < len(seats); i++ {
		sum += int(math.Abs(float64(seats[i] - students[i])))
	}
	return sum
}

type Bank struct {
	account []int
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	n := len(balance)
	account := make([]int, n+1)
	balance1 := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		account[i] = i
		balance1[i] = balance[i-1]
	}
	return Bank{
		account: account,
		balance: balance1,
		n:       n + 1,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.balance[account1] < money || account1 >= this.n || account2 >= this.n {
		return false
	}
	this.balance[account1] -= money
	this.balance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account >= this.n {
		return false
	}
	this.balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account >= this.n {
		return false
	}
	if this.balance[account] < money {
		return false
	}
	this.balance[account] -= money
	return true
}

func countMaxOrSubsets(nums []int) int {
	arr := []int{}
	numMap := make(map[int]int)
	max := 0
	var helper func(nums, arr []int, index int)
	helper = func(nums, arr []int, index int) {
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum = sum | arr[i]
		}
		numMap[sum]++
		if sum > max {
			max = sum
		}
		for i := index; i < len(nums); i++ {
			arr = append(arr, nums[i])
			helper(nums, arr, i+1)
			arr = arr[:len(arr)-1]
		}
	}
	helper(nums, arr, 0)
	return numMap[max]
}
