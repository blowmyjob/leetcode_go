package leetcode

type StockSpanner struct {
	stack    []int
	minStack []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.stack = append(this.stack, price)
	for len(this.minStack) > 0 && this.stack[this.minStack[len(this.minStack)-1]] <= price {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	count := 0
	if len(this.minStack) == 0 {
		count = len(this.stack)
	} else {
		count = len(this.stack) - this.minStack[len(this.minStack)-1] - 1
	}
	this.minStack = append(this.minStack, len(this.stack)-1)
	return count
}
