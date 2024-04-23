package best_time_to_buy_and_sell_stock_ii

func MaxProfit(prices []int) int {
	var mP int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			mP += prices[i] - prices[i-1]
		}
	}
	return mP
}
