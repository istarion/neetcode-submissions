func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	suffix := 1
	zeros := 0
	zeropos := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros++
			zeropos = i
			continue
		}
		suffix *= nums[i]
	}
	switch {
		case zeros > 1: return result
		case zeros == 1:
			result[zeropos] = suffix
			return result
	}
	prefix := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = prefix * suffix
			prefix *= nums[i]
		} else {
			suffix /= nums[i]
			result[i] = prefix * suffix
			prefix *= nums[i]
		}
	}
	return result
}