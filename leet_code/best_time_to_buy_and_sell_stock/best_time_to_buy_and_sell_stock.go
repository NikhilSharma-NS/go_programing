package best_time_to_buy_and_sell_stock

import "fmt"

func MaxProfit(prices []int) int {
	max := 0
	index := 0
	start := 1

	for i := 0; i < len(prices) && start < len(prices); i++ {
		value := prices[start] - prices[i]
		if value > 0 && value > max {
			index = start
			max = value
		}
		start++

	}
	if max == 0 {
		fmt.Println("prices:max", max)
		return 0
	}

	fmt.Println("prices", "", max)
	return prices[index]
}
