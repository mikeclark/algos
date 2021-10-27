package simplesearch

func search(search int, theArray []int) int {
	for i, val := range theArray {
		if search == val {
			return i
		}
	}

	// no good
	return -1
}
