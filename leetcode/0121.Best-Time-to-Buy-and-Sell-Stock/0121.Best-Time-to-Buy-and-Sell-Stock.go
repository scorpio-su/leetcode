package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	price, fit := 1000000, 0
	for _, p := range prices {
		price = min(p, price)
		fit = max(fit, p-price)
	}
	return fit
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
