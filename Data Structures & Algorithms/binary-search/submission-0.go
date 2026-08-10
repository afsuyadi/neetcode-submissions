
func search(nums []int, target int) int {
	// find target inside nums
	// if found, return the index it currently loops
	for i, val := range nums {
		if val == target {
			return i
		} else {
			continue
		}
	}
	return -1
}