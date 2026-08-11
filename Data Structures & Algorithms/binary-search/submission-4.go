
func search(nums []int, target int) int {
	leftIdx := 0
	rightIdx := len(nums)-1
	for leftIdx <= rightIdx {
		midIdx := (leftIdx + rightIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		} else if target < nums[midIdx] {
			rightIdx = midIdx - 1
		} else {
			leftIdx = midIdx + 1
		}
	}
	return -1
} 