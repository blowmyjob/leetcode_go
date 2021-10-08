package leetcode

func removeDuplicateLetters(s string) string {
	// 每个字符出现次数
	var count [26]int
	// 是否在栈中，存在为true
	var exist [26]bool
	// 单调栈
	var stack []rune

	// 统计字符出现次数
	// 注：c-'a'的结果是数字，即距离字母'a'有多远，比如'a'-'a'=0, 即就是a本身
	for _, c := range s {
		count[c-'a']++
	}
	for _, c := range s {
		if exist[c-'a'] {
			count[c-'a']--
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 {
			exist[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		count[c-'a']--
		exist[c-'a'] = false
	}
	return string(stack)
}
