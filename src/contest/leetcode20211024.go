package contest

import (
	"math"
	"strconv"
	"strings"
)

func NextBeautifulNumber(n int) int {
	var jungde func(n int) bool
	jungde = func(n int) bool {
		str := strconv.Itoa(n)
		for i := 0; i <= 9; i++ {
			count := strings.Count(str, strconv.Itoa(i))
			if count != 0 && count != i {
				return false
			}
		}
		return true
	}
	strN := strconv.Itoa(n)
	max := int(math.Pow(10, float64(len(strN)+1)))
	for i := n + 1; i <= max; i++ {
		if jungde(i) {
			return i
		}
	}
	return 0
}
