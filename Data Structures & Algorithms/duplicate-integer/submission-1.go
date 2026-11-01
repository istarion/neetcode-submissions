func hasDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
