package leetcode

func generateParenthesis(n int) []string {
	s := make([]string, 0, 0)
	t := ""
	funcW(&s, t[:0], 0, 0)
	return s
}

func funcW(s *[]string, t string, l int, r int) {
	if l == 0 && r == 0 {
		*s = append(*s, t)
	}
	if l > r || l < 0 || r < 0 {
		return
	}
	funcW(s, t+"(", l-1, r)
	funcW(s, t+")", l, r-1)

}
