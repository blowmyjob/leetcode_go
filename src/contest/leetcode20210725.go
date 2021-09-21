package contest

func MaximumNumber(num string, change []int) string {
	s := []byte(num)
	i, n := 0, len(num)
	// 从高到低，找到第一个映射后比现在大的位置
	for ; i < n; i++ {
		old := int(s[i] - '0')
		new := change[old]
		if old < new {
			break
		}
	}
	// 从高到底，开始映射，直到映射后比现在小
	for ; i < n; i++ {
		old := int(s[i] - '0')
		new := change[old]
		if old > new {
			break
		}
		s[i] = '0' + byte(new)

	}
	return string(s)
}
