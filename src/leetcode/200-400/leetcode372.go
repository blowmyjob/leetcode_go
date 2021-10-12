package leetcode

func superPow(a int, b []int) int {
	base := 1337
	size := len(b)
	if size == 0 {
		return 1
	}
	var pow func(source, index int) int
	pow = func(source, index int) int {
		if index == 0 {
			return 1
		}
		//奇数幂
		if index&1 == 1 {
			return (source % base * pow(source, index-1)) % base
		} else {
			//偶数幂
			sub := pow(source, index/2)
			return (sub * sub) % base
		}
	}
	last := b[size-1]
	b = b[:size-1]
	return (pow(a, last) * pow(superPow(a, b), 10)) % base
}
