
func maxProfit(prices []int) int {
	// i iteration will always going forward
	// if the biggest value is in i = 0, then always return 0 (cant sell it)

	// brute-force
	currentStored := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] - prices[i] > 0 { // if it has profit
				if currentStored < prices[j] - prices[i] {
					currentStored = prices[j] - prices[i]
				} else {
					continue
				}
			} else { // either it is a loss or 0
				continue // skip this comparison
			}
		}
	}
	return currentStored
}