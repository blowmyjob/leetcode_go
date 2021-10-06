package leetcode

func majorityElement(nums []int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}
	cand1, count1 := nums[0], 0
	cand2, count2 := nums[0], 0
	for _, v := range nums {
		if cand1 == v {
			count1++
			continue
		}
		if cand2 == v {
			count2++
			continue
		}
		if count1 == 0 {
			cand1 = v
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = v
			count2++
			continue
		}
		count1--
		count2--
	}
	count1 = 0
	count2 = 0
	for _, v := range nums {
		if cand1 == v {
			count1++
		} else if cand2 == v {
			count2++
		}
	}
	if count1*3 > len(nums) {
		res = append(res, cand1)
	}
	if count2*3 > len(nums) {
		res = append(res, cand2)
	}
	return res
}
