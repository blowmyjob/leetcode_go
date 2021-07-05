package contest

func FindingUsersActiveMinutes(logs [][]int, k int) []int {
	nums := killRepetion(logs)
	users := make(map[int]int)
	count := make([]int, k)
	for i := 0; i < len(nums); i++ {
		key := nums[i][0]
		users[key]++
	}
	for _, value := range users {
		count = append(count, value)
	}
	res := make([]int, k)
	for i := 0; i < len(count); i++ {
		if count[i] <= k && count[i] >= 1 {
			res[count[i]-1]++
		}
	}
	return res
}

func killRepetion(nums [][]int) [][]int {
	newRes := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[i][0] == nums[j][0] && nums[i][1] == nums[j][1] {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes
}

func purchasePlans(nums []int, target int) int {
	num := 0

	quickSort(nums, 0, len(nums))
	start, end := 0, len(nums)-1
	for start < end {
		t := nums[start] + nums[end]
		if t <= target {
			num++
			start++
		} else {
			end--
		}
	}
	return num % 1000000007
}

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func StoreWater(bucket []int, vat []int) int {
	var sum = 0
	for k, _ := range bucket {
		var min = 0
		if bucket[k] >= vat[k] {
			min = -1
		} else if bucket[k] == 0 {
			continue
		} else {
			min = vat[k] / bucket[k]
		}
		if min < sum {
			sum = min
		}
	}
	return 0
}
