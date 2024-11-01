package besttime

func MaxProfit(price []int) int {
	minPrice := price[0]
	maxProfit := 0

	for _, priceOfDay := range price[1:] {
		//находим минимум
		if priceOfDay < minPrice {
			minPrice = priceOfDay
		}
		//находим профит
		if maxProfit < priceOfDay-minPrice {
			maxProfit = priceOfDay - minPrice
		}
	}
	return maxProfit
}
