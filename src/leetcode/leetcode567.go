package leetcode

func checkInclusion(s1 string, s2 string) bool {
	count1, count2 := [26]int{}, [26]int{} // 两个长度为26的数组
	for i := 0; i < len(s1); i++ {         // 统计s1的字符的出现次数
		count1[s1[i]-'a']++
	}
	start := 0             // 指向s2子串的开头
	for i, c := range s2 { // 遍历s2的字符
		count2[c-'a']++ // 将s2的字符统计到count2

		// 开启循环，在start没有越界的前提下，start上的字符，count2的比count1多
		// 这说明start不是合格子串的开头，start应该步进，舍弃这个字符
		for start <= i && count1[s2[start]-'a'] < count2[s2[start]-'a'] {
			count2[s2[start]-'a']--
			start++
		}
		// 如果两个count数组全等，说明找到了满足条件的子串
		if count1 == count2 { // go里面数组是基本类型，可以这么比较
			return true
		}
	}
	// 考察了所有s2的字符，都没有找到合适的子串
	return false
}
