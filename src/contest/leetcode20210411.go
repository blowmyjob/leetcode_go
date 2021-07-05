package contest

func ArraySign(nums []int) int {
	var sum = 1
	for _, v := range nums {
		if v > 0 {
			sum *= 1
		} else if v < 0 {
			sum *= -1
		}
		if v == 0 {
			return 0
		}
	}
	if sum > 0 {
		return 1
	} else {
		return -1
	}
}

func findTheWinner(n int, k int) int {
	var find, count int
	num := 0        //当num到达N-1时，只剩一只猴子
	flag := []int{} //flag[i],该猴子已经出去了

	for i := 0; i < n; i++ {
		flag = append(flag, 0)
	}
	i := 0 //找下一只要出去的猴子从i开始找
	for num != n-1 {
		find = 0
		count = 0
		for ; find == 0; i++ {
			t := i % n
			if flag[t] == 0 {
				count += 1
			}
			if count == k {
				num++
				find = 1
				flag[t] = 1
			}
		}
	}
	for k, v := range flag {
		if v == 0 {
			return k + 1
		}
	}
	return 0
}
