package leetcode

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	arrs := make([]int, len(nums)+1)
	for k, v := range nums {
		if k == 0 {
			arrs[k] = v
		} else {
			arrs[k] = arrs[k-1] + v
		}
	}
	return NumArray{
		arr: arrs,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.arr[right]
	}
	return this.arr[right] - this.arr[left-1]
}
