package quicksort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	assertSortedArray := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Testing Sort 1", func(t *testing.T) {
		testArray := []int{8, 1, 45, 6, 67, 7, 82, 89, 2}
		got := Sort(testArray)
		want := []int{1, 2, 6, 7, 8, 45, 67, 82, 89}
		assertSortedArray(t, got, want)
	})
}
