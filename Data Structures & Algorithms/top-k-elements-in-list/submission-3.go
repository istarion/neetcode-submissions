type ctrArray [2001]int

func topKFrequent(nums []int, k int) []int {
	var ctr ctrArray
	for _, n := range nums {
		ctr[n+1000]++
	}
	result := make([]int, k)

	kmax := getKMaxPos(ctr, k)
	for ik, pos := range kmax {
		result[ik] = pos - 1000
	}
	return result
}

type kMax struct {
	max []int
	pos []int
}

func newKMax(k int) *kMax {
	max := make([]int, k)
	pos := make([]int, k)
	for i := range max {
		max[i] = math.MinInt
		pos[i] = -1
	}
	return &kMax{
		max: max,
		pos: pos,
	}
}

func (km *kMax) add(n int, p int) {
	km.max[0] = n
	km.pos[0] = p
	for i := 1; i < len(km.max); i++ {
		if km.max[i] < km.max[i-1] {
			km.max[i], km.max[i-1] = km.max[i-1], km.max[i]
			km.pos[i], km.pos[i-1] = km.pos[i-1], km.pos[i]
		} else {
			break
		}
	}
}

func getKMaxPos(nums ctrArray, k int) []int {
	kmax := newKMax(k)
	for i, n := range nums {
		if n > kmax.max[0] {
			kmax.add(n, i)
		}
	}
	return kmax.pos
}
