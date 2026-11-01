type anagramKey [26]int
type anagramGroup = map[anagramKey][]string

func toAnagramKey(s string) anagramKey {
	ak := anagramKey{}
	for _, c := range s {
		ak[c-'a']++
	}
	return ak
}

func groupAnagrams(strs []string) [][]string {
	ag := make(anagramGroup)
	for _, str := range strs {
		ak := toAnagramKey(str)
		if _, exists := ag[ak]; !exists {
			ag[ak] = []string{}
		}
		ag[ak] = append(ag[ak], str)
	}
	result := make([][]string, 0, len(ag))
	for _, group := range ag {
		result = append(result, group)
	}
	return result
}