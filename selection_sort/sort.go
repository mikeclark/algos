package selection_sort

func Sort(arr []int) []int {
	num := len(arr)
	for x := 0; x < num; x++ {
		first := x
		for next := x; next < num; next++ {
			if arr[next] < arr[first] {
				first = next
			}
		}
		arr[x], arr[first] = arr[first], arr[x]
	}
	return arr
}
