package contest

import "sort"

func FindFarmland(land [][]int) [][]int {
	rowlen, colLen := len(land), len(land[0])
	res := [][]int{}
	for i := 0; i < rowlen; i++ {
		for j := 0; j < colLen; j++ {
			if land[i][j] == 1 {
				maxRow := i
				maxCol := j
				for row := i + 1; row < rowlen && land[row][j] == 1; row++ {
					maxRow++
				}
				for col := j + 1; col < colLen && land[i][col] == 1; col++ {
					maxCol++
				}
				for row := i; row <= maxRow; row++ {
					for col := j; col <= maxCol; col++ {
						land[row][col] = 0
					}
				}
				res = append(res, []int{i, j, maxRow, maxCol})
			}
		}
	}
	return res
}

type LockingTree struct {
	parent   []int
	treeNode [][]int
	user     []int
}

func Constructor2(parent []int) LockingTree {
	n := len(parent)
	treeNode := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parent[w]
		treeNode[v] = append(treeNode[v], w)
	}
	return LockingTree{
		parent:   parent,
		user:     make([]int, n),
		treeNode: treeNode,
	}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.user[num] > 0 {
		return false
	}
	this.user[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.user[num] != user {
		return false
	}
	this.user[num] = 0
	return true
}

// 判断 v 的子孙是否有锁
func (this *LockingTree) hasLock(v int) bool {
	for _, w := range this.treeNode[v] {
		if this.user[w] > 0 || this.hasLock(w) {
			return true
		}
	}
	return false
}

// 解锁 v 的所有子孙
func (this *LockingTree) unlock(v int) {
	for _, w := range this.treeNode[v] {
		this.user[w] = 0
		this.unlock(w)
	}
}

func (this *LockingTree) Upgrade(num, user int) bool {
	for v := num; v >= 0; v = this.parent[v] {
		if this.user[v] > 0 {
			return false
		}
	}
	if !this.hasLock(num) {
		return false
	}
	this.user[num] = user
	this.unlock(num)
	return true
}

func NumberOfWeakCharacters(properties [][]int) int {
	ans := 0
	sort.Slice(properties, func(i, j int) bool {
		a, b := properties[i], properties[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return ans
}
