package main

func findKeyCount(data []int, key int) int {
	count := 0
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == key {
			count++
		}
	}

	return count
}
