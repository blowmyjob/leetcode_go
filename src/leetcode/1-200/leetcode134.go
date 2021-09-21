package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	allSave, curSave, position := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		allSave += gas[i] - cost[i]
		curSave += gas[i] - cost[i]
		if curSave < 0 {
			position = i + 1
			curSave = 0
		}
	}
	if allSave >= 0 {
		return position
	} else {
		return -1
	}
}
