package offer

import "strings"

func ReplaceSpace(s string) string {
	//write your code here
	s = strings.ReplaceAll(s, " ", "%20")
	return s
}
