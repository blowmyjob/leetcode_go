package leetcode

func LongestConsecutive(nums []int) int {
	res := 0
	wordMap := make(map[int]int)
	for _, v := range nums {
		wordMap[v] = 1
	}
	for _, v := range nums {
		if wordMap[v] == 1 {
			num := 1
			curNum := v - 1
			for wordMap[curNum] == 1 {
				delete(wordMap, curNum)
				curNum--
				num++
			}
			curNum = v + 1
			for wordMap[curNum] == 1 {
				delete(wordMap, curNum)
				curNum++
				num++
			}
			if num > res {
				res = num
			}
			delete(wordMap, v)
		}
	}
	return res
}
