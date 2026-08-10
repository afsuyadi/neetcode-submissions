func twoSum(numbers []int, target int) []int {
	// try comparing the front and back
	// result := make([]int, 2)
	// for i := 1; i < len(numbers); i++ {
	// 	for j := i+1; j < len(numbers); j++ {
	// 		summed := i + j
	// 		if summed == target &&  i < j {
	// 			result[0] = i
	// 			result[1] = j
	// 		}
	// 	} 
	// }
	// return result
	// TLE
	left := 0 // ex. 1
	right := len(numbers) - 1 // ex. 4
	
	for left < right {
		summed := numbers[left] + numbers[right]
		if summed == target && left != right { // 1+ 4 != 3
			return []int{left+1, right+1}
		} else if summed > target {
			right--
		} else {
			left++
		}
	}
	return []int{left, right}	
}
