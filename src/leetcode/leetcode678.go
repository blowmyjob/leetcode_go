package leetcode

func checkValidString(s string) bool {
	leftStack := []int{}
	starStack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftStack = append(leftStack, i)
		} else if s[i] == '*' {
			starStack = append(starStack, i)
		} else {
			if len(leftStack) == 0 {
				if len(starStack) == 0 {
					return false
				} else {
					starStack = starStack[:len(starStack)-1]
				}
			} else {
				leftStack = leftStack[:len(leftStack)-1]
			}
		}
	}
	if len(leftStack) > len(starStack) {
		return false
	}
	for len(leftStack) != 0 && len(starStack) != 0 {
		if leftStack[len(leftStack)-1] > starStack[len(starStack)-1] {
			return false
		}
		leftStack = leftStack[:len(leftStack)-1]
		starStack = starStack[:len(starStack)-1]
	}
	return len(leftStack) == 0
}
