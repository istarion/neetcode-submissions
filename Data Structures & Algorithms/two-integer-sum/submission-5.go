func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, exists := m[num]; exists && num*2 == target {
			return []int{j, i}
		}
		m[num] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok && i != j {
			return []int{i, j}
		}
	}
	panic("no solution")
}
