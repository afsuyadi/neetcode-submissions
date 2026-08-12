
func minEatingSpeed(piles []int, h int) int {
	getTotalPilesHour := func (midK int) int {
		result := 0 // banana
		// midK unit is banana per hour
		for _, pile := range piles {
			result += ( (pile + midK - 1 )/ midK )
		}
		return result // return number of hours
	}

	getHighestK := func(bananas []int) int {
		biggestPile := piles[0] // ex. 25
		for _, pile := range bananas { // ex. 5
			if biggestPile < pile {
				biggestPile = pile
			}
		}
		return biggestPile
	}
	// k max will produce lowest h (hour) by using the max value of pile
	// we use the range value of k low and k high as space of search
	// we use binary search
	// while binary search is conducted, we use internal function to get total of piles in that current k
	lowK := 1 // lowest k always start from 1
	highK := getHighestK(piles)// get highest k 
	answer := highK // use a 4th var as the memory to remember mid K
	for lowK <= highK {
		midK := (lowK + highK) / 2 // set midK as starting point
		isStillLower := getTotalPilesHour(midK) <= h // use True or False to check if midTotalHour IS STILL LOWER than target h hour.
		if isStillLower { // means that MidTotalHour is still lower than h hour, which is good
			highK = midK - 1 // increase the lower K because it is still safe
			answer = midK // remember midK, in case in the next iteration, midK produces HIGHER total h
		} else { // means that midTotalHour is now higher than h hour, which is bad
			lowK = midK + 1 // increase lowK
		}
	}
	return answer // means that the highest k will product the lowest h.
}





// You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

// You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.

// Return the minimum integer k such that you can eat all the bananas within h hours.