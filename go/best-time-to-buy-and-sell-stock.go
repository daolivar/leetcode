package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var start, maxProfit int
	for end := 1; end < len(prices); end++ {
		if prices[start] < prices[end] {
			maxProfit = max(maxProfit, prices[end]-prices[start])
		} else {
			start = end
		}
	}
	return maxProfit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
