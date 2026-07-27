func topKFrequent(nums []int, k int) []int {
	// map setiap elemen di dalam nums terhadap freq nya
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	// output = {
		// 1: 1,
		// 2: 2,
		// 3: 3,	
	//}
	// reorder element memakai freq sebagai index nya 
	buckets := make([][]int, len(nums)+1) // eg. [[],[],[],[],[],[],[]]
	for num, freq := range count { // eg. 3:3
		buckets[freq] = append(buckets[freq], num) // eg. [[],[],[],[3],[],[],[]]
	}
	// output= [[],[1],[2],[3]],[],[],[]]
	// mulai dari belakang
	result := []int{} 
	for freq := len(buckets) - 1; freq > 0; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
			
	}
	return []int{}
}
