package binarysearch

func search(search int, theArray []int) int {
	var begin = 0
	var end = len(theArray) - 1
	for begin <= end {
		// find halway
		var mid = int((begin + end) / 2)
		var val = theArray[mid]

		// check val
		if val == search {
			return mid
		} else if val < search {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
