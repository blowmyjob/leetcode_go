package leetcode

import "path/filepath"

func simplifyPath(path string) string {
	return filepath.Clean(path)
}
