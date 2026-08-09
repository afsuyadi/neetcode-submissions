type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// decoded_string := make([]string, len(strs))
	var decoded_string string
	for _, val := range strs {
		val = val + string('😂')
		decoded_string += val
	}
	return decoded_string
}

func (s *Solution) Decode(encoded string) []string {
	var finalword strings.Builder
	result := make([]string, 0)
	for _, letter := range encoded {
		if letter == '😂' {
			result = append(result, finalword.String())
			finalword.Reset()
		} else {
			finalword.WriteRune(letter)
		}
		
	}
	return result
}
