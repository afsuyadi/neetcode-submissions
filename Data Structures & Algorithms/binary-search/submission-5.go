func search(nums []int, target int) int {
	leftIdx := 0
	rightIdx := len(nums)-1
	for leftIdx <= rightIdx {
		midIdx := (rightIdx + leftIdx) / 2
		if nums[midIdx] > target {
			rightIdx = midIdx - 1
		} else if nums[midIdx] < target {
			leftIdx = midIdx + 1
		} else {
			return midIdx
		}
	}
	return -1
}