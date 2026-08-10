func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	zeroCount := 0
	totalMultiply := 1
	zeroIdx := -1

	for i, val := range nums {
		if val == 0 {
			zeroCount ++
			zeroIdx = i	
		} else {
			totalMultiply = totalMultiply * val
		}
		
	}
	if zeroCount == 1 {
		result[zeroIdx] = totalMultiply
		return result
	}
	if zeroCount > 1 {
		return result
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
		
		current = totalMultiply / nums[i]
		 
		// else if nums[i] == 0 {
		// 	current = totalMultiply
		// } else if nums[i] != 0 && zeroCount != 0 {
		// 	current = 0
		// }
		
		result[i] = current
	}
	return result
}
