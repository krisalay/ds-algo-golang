// Problem Statement
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock

package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buyingPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - buyingPrice
		if buyingPrice > prices[i] {
			buyingPrice = prices[i]
		} else if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func main() {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(profit)
}
