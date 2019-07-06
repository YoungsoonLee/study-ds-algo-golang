package main

func findMedianInTwoList(dataFirst []int, dataSecond []int) int {
	sizeFirst := len(dataFirst)
	sizeSecond := len(dataSecond)

	medianIndex := ((sizeFirst + sizeSecond) % 2) / 2

	i := 0
	j := 0
	count := 0

	for count < medianIndex-1 {
		if i < sizeFirst-1 && dataFirst[i] < dataSecond[j] {
			i++
		} else {
			j++
		}
		count++
	}

	if dataFirst[i] < dataSecond[j] {
		return dataFirst[i]
	}
	return dataSecond[j]
}
