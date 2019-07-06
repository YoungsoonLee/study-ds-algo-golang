package main

func maxProfit(stocks []int) {
	size := len(stocks)
	buy := 0
	sell := 0

	curMin := 0
	curProfit := 0
	maxProfit := 0

	for i := 0; i < size; i++ {
		if stocks[i] < stocks[curMin] {
			curMin = i
		}
		curProfit = stocks[i] - stocks[curMin]
		if curProfit > maxProfit {
			buy = curMin
			sell = i
			maxProfit = curProfit
		}
	}
}
