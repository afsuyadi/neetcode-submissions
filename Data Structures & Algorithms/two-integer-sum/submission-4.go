func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if index, ok := seen[diff]; ok {
			return []int{index, i}
		} else {
			seen[val] = i
		}
	} 
	return []int{}
}
