func productExceptSelf(nums []int) []int {
	// we loop through every element in nums
	// we multiply the value of other nums, except itself
	// we append into the new array
	// ex. nums = [1,2,4,6]
	var result []int
	zeroCount := 0
	for _, val := range nums {
		if val == 0 {
			zeroCount ++
		}
	}
	if zeroCount > 1 {
		allZero := make([]int, len(nums))
		return allZero
	}

	// for i := 0; i < len(nums); i++ { // len(nums) = 4
	// 	// ex: i = 2, nums[2] = 4
	// 	current := 1
	// 	if nums[i] == 0 {
	// 		continue
	// 	}
	// 	for j := 0; j < len(nums); j++ {
	// 		// ex. i = 0, 
	// 		// j = 1, nums[1] = 2,
	// 		// j = 2, nums[2] = 4,
	// 		// j = 3, nums[3] = 6,
	// 		if j == i {
	// 			continue
	// 		}
	// 		if nums[j] == 0 {
	// 		continue
	// 	}
			
	// 		current = current * nums[j]
	// 	}		
	// 	result = append(result, current)
	// }
	totalMultiply := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			totalMultiply = totalMultiply * nums[i]
		}
		continue		
	}
	for i := 0; i < len(nums); i++ { // len(nums) = 4
		// ex: i = 2, nums[2] = 4
		current := 1
		if nums[i] != 0 && zeroCount == 0 { // does not have any 0 AND the value is NOT zero
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
