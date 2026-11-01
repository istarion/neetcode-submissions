type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		for _, c := range str {
			if c == '\\' {
				sb.WriteByte('\\')
				sb.WriteByte('\\')
			} else {
				sb.WriteRune(c)
			}
		}
		sb.WriteByte('\\')
		sb.WriteByte('n')
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)
	escaped := false
	var sb strings.Builder
	for _, c := range encoded {
		if escaped {
			if c == 'n' {
				result = append(result, sb.String())
				sb.Reset()
			} else {
				sb.WriteRune(c)
			}
			escaped = false
		} else if c == '\\' {
			escaped = true
		} else {
			sb.WriteRune(c)
		}
	}
	return result
}