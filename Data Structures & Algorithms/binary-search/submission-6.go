// func search(nums []int, target int) int {
// 	leftIdx := 0
// 	rightIdx := len(nums)-1
// 	for leftIdx <= rightIdx {
// 		midIdx := (rightIdx + leftIdx) / 2
// 		if nums[midIdx] > target {
// 			rightIdx = midIdx - 1
// 		} else if nums[midIdx] < target {
// 			leftIdx = midIdx + 1
// 		} else {
// 			return midIdx
// 		}
// 	}
// 	return -1
// }

func binarySearchRecursive(nums []int, target int, leftIdx int, rightIdx int) int {
	if leftIdx > rightIdx {
		return -1
	}
	midIdx := (leftIdx + rightIdx) / 2
	if nums[midIdx] > target {
		rightIdx = midIdx - 1
	} else if nums[midIdx] < target {
		leftIdx = midIdx + 1
	} else {
		return midIdx
	}
	return binarySearchRecursive(nums, target, leftIdx, rightIdx)
}
func search(nums []int, target int) int {
	return binarySearchRecursive(nums, target, 0, len(nums)-1)
}