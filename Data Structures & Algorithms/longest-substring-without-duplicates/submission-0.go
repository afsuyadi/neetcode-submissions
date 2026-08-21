
func checkDuplicate(hashmap map[string]int, key string) bool {
	if _, exists := hashmap[key]; exists {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	
	leftPtr := 0
	rightPtr := 0
	hashmap := make(map[string]int)
	maxCount := 0

	// ====v1====
	// for rightPtr := 0; rightPtr < len(s); rightPtr++ {
	// 	if _, exist := hashmap[string(s[rightPtr])]; !exist {
	// 		// if it doesnt exist?
	// 		hashmap[string(s[rightPtr])] += 1
	// 		// continue, so rightPtr can move + 1
	// 	} else { // if it exists (which means it's duplicated)
	// 		leftPtr += 1 // left pointer move +1
	// 	}
	// 	maxCount = rightPtr - leftPtr
	// }
	// return maxCount
	
	// ==== v2 ====
	for leftPtr <= rightPtr && rightPtr < len(s) {
		rightKey := string(s[rightPtr])
		leftKey := string(s[leftPtr])
		// isDuplicate := checkDuplicate(hashmap, rightKey)	
		for checkDuplicate(hashmap, rightKey) {
			// delete(hashmap, rightKey)
			leftKey = string(s[leftPtr])
			delete(hashmap, leftKey)
			leftPtr += 1
		}

		rightPtr += 1
		hashmap[rightKey] = 1
		// if isDuplicate == true {
		// 	leftPtr += 1
		// 	delete(hashmap, leftKey)
		// } else {
		// 	rightPtr += 1
		// 	hashmap[rightKey] = 1
		// }
		
		windowLen := rightPtr - leftPtr
		if windowLen > maxCount {
			maxCount = windowLen
		}
		// rightPtr += 1
	}
	fmt.Println(hashmap)
	return maxCount
}

