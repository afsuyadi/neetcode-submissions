
func getRightRow(m [][]int, target int) int {
	lowRowIdx := 0 // the lower row
	highRowIdx := len(m)-1 // the higher row
	for lowRowIdx <= highRowIdx {
		midRowIdx := (lowRowIdx + highRowIdx)/2
		if m[midRowIdx][len(m[midRowIdx])-1] < target {
			lowRowIdx = midRowIdx + 1
		} else if m[midRowIdx][0] > target {
			highRowIdx = midRowIdx - 1
		} else {
			return midRowIdx
		}
	}
	return -1
}

func isTargetValid(rows []int, target int) bool {
	lowIdx := 0
	highIdx := len(rows)-1
	
	for lowIdx <= highIdx {
		midIdx := (lowIdx + highIdx)/2
		if rows[midIdx] < target {
			lowIdx = midIdx + 1
		} else if rows[midIdx] > target {
			highIdx = midIdx - 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {

	rightRow := getRightRow(matrix, target)
	if rightRow != -1 {
		return isTargetValid(matrix[rightRow], target)
	} else { return false }
}

// You are given an m x n 2-D integer array matrix and an integer target.

// Each row in matrix is sorted in non-decreasing order.
// The first integer of every row is greater than the last integer of the previous row.
// Return true if target exists within matrix or false otherwise.

// Can you write a solution that runs in O(log(m * n)) time?