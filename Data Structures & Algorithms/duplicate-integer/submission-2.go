func hasDuplicate(nums []int) bool { // input nums with array typed integers. return bool
    seen := make(map[int]bool)
    for _, x := range nums {
        if seen[x] {
            return true
        } else {
            seen[x] = true
        }
    }
    return false
}
