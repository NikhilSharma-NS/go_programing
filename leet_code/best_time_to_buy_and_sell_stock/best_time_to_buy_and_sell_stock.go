package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	max := 0

	for buy, sell := 0, 1; sell < len(prices); {
		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			if profit > max {
				max = profit
			}
		} else {
			buy = sell
		}
		sell = sell + 1

	}

	return max
}
