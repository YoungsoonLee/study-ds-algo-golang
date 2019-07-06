package main

func find01List(data []int) int {
	size := len(data)
	if size == 1 && data[0] == 1 {
		return -1
	}

	return find01ListUntil(data, 0, size-1)
}

func find01ListUntil(data []int, start int, end int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2

	if 1 == data[mid] && 0 == data[mid-1] {
		return mid
	}

	if 0 == data[mid] {
		return find01ListUntil(data, mid+1, end)
	}
	return find01ListUntil(data, start, mid-1)

}
