type ctrArray [2001]int

func topKFrequent(nums []int, k int) []int {
	var ctr ctrArray
	for _, n := range nums {
		ctr[n+1000]++
	}
	result := make([]int, k)
	for ik := 0; ik < k; ik++ {
		maxPos := getMaxPos(ctr)
		result[ik] = maxPos - 1000
		ctr[maxPos] = math.MinInt
	}
	return result
}

func getMaxPos(nums ctrArray) int {
	max := math.MinInt
	maxPos := -1
	for i, n := range nums {
		if n > max {
			max = n
			maxPos = i
		}
	}
	return maxPos
}