
func search(nums []int, target int) int {
	// find target inside nums
	// if found, return the index it currently loops

	// non-Binary Search
	// for i, val := range nums {
	// 	if val == target {
	// 		return i
	// 	} else {
	// 		continue
	// 	}
	// }
	// return -1

	// Binary Search
	low, high := 0, len(nums)-1	 // 0, 6
	mid := (low + high) / 2
	if target < nums[low] || target > nums[high] {
		return -1
	}
	for low <= high {
		
		// if high % 2 != 0 { // odd number
		// 	mid = ((high + 1 - low) / 2) - 1 // mid index for odd high
		// } else {
		// 	mid = (high / 2) - 1
		// }

		
		if target > nums[mid] {
			mid += 1
			low += 1
		} else if target < nums[mid] {
			mid -= 1
			high -= 1
		} else if target == nums[mid] {
			return mid
		} else {
			return -1
		}
	}
	return -1

}
