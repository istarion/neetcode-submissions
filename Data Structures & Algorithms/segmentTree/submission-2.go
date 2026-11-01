type SegmentTree struct {
	sum int
	l   int
	r   int
	lc  *SegmentTree
	rc  *SegmentTree
}

func NewSegmentTree(nums []int) *SegmentTree {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l int, r int) *SegmentTree {
	if l == r {
		return &SegmentTree{
			sum: nums[l], l: l, r: r,
		}
	}
	m := (l + r) / 2
	lc := build(nums, l, m)
	rc := build(nums, m+1, r)
	return &SegmentTree{
		sum: lc.sum + rc.sum, l: l, r: r, lc: lc, rc: rc,
	}
}

func (st *SegmentTree) Update(index, val int) {
	st.updateHelper(st, index, val)
}

func (st *SegmentTree) updateHelper(cur *SegmentTree, index, val int) {
	if cur.l == cur.r {
		cur.sum = val
		return
	}
	m := (cur.l + cur.r) / 2
	if index > m {
		st.updateHelper(cur.rc, index, val)
	} else {
		st.updateHelper(cur.lc, index, val)
	}
	cur.sum = cur.lc.sum + cur.rc.sum
	return
}

func (st *SegmentTree) Query(L, R int) int {
	return st.queryHelper(st, L, R)
}

func (st *SegmentTree) queryHelper(cur *SegmentTree, L, R int) int {
	if cur.r < L || cur.l > R {
		return 0
	}
	if cur.l >= L && cur.r <= R {
		return cur.sum
	}
	return st.queryHelper(cur.lc, L, R) + st.queryHelper(cur.rc, L, R)
}