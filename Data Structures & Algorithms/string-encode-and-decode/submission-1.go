type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	data, _:= json.Marshal(strs)
	return string(data)
}

func (s *Solution) Decode(encoded string) []string {
	var strs []string
	json.Unmarshal([]byte(encoded), &strs)

	return strs
}
