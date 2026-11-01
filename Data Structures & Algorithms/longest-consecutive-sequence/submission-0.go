func longestConsecutive(nums []int) int {
	values := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		values[num] = struct{}{}
	}

	maxLength := 0
	for num := range values {
		if _, exists := values[num-1]; exists {
			continue
		}

		length := 1
		for current := num + 1; ; current++ {
			if _, exists := values[current]; !exists {
				break
			}
			length++
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}