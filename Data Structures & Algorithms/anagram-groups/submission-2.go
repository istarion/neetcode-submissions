type anagramKey [26]int
type anagramGroup = map[anagramKey][]int

func toAnagramKey(s string) anagramKey {
	ak := anagramKey{}
	for _, c := range s {
		ak[c-'a']++
	}
	return ak
}

func groupAnagrams(strs []string) [][]string {
	ag := make(anagramGroup)
	for i, str := range strs {
		ak := toAnagramKey(str)
		if _, exists := ag[ak]; !exists {
			ag[ak] = make([]int, 0)
		}
		ag[ak] = append(ag[ak], i)
	}
	result := make([][]string, 0, len(ag))
	for _, group := range ag {
		anagrams := make([]string, 0, len(group))
		for _, n := range group {
			anagrams = append(anagrams, strs[n])
		}
		result = append(result, anagrams)
	}
	return result
}