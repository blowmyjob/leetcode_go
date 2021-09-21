package contest

import "sort"

func SmallestChair(times [][]int, targetFriend int) int {
	targetTime := times[targetFriend]
	chairs := make([]int, len(times))
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	for i := 0; i < len(times); i++ {
		tmp := times[i]
		for j := 0; j < len(chairs); j++ {
			if chairs[j] <= tmp[0] {
				chairs[j] = tmp[1]
				if targetTime[0] == tmp[0] && targetTime[1] == tmp[1] {
					return j
				}
				break
			}
		}
	}
	return -1
}
