package quicksort

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// set less/greater than pivot vars
	l := make([]int, 0, len(arr))
	g := make([]int, 0, len(arr))

	// selecting first element for pivot...could lead
	// to worst case scenrios
	pivot := arr[0]

	for i := range arr {
		if i != 0 {
			if arr[i] < pivot {
				l = append(l, arr[i])
			} else {
				g = append(g, arr[i])
			}
		}
	}
	arr = nil

	return append(append(Sort(l), pivot), Sort(g)...)
}
