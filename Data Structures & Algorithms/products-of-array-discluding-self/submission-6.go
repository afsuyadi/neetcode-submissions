func productExceptSelf(nums []int) []int {
	var result []int
	zeroCount := 0
	totalMultiply := 1
	for _, val := range nums {
		if val == 0 {
			zeroCount ++
		} else {
			totalMultiply = totalMultiply * val
		}
		
	}
	if zeroCount > 1 {
		allZero := make([]int, len(nums))
		return allZero
	}

	// totalMultiply := 1
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != 0 {
	// 		totalMultiply = totalMultiply * nums[i]
	// 	}
	// 	continue		
	// }
	for i := 0; i < len(nums); i++ {
		current := 1
		if nums[i] != 0 && zeroCount == 0 {
			current = totalMultiply / nums[i]
		} else if nums[i] == 0 {
			current = totalMultiply
		} else if nums[i] != 0 && zeroCount != 0 {
			current = 0
		}
		
		result = append(result, current)
	}
	return result
}
