package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	// t.Fatal("not implemented")
	testArray := []int{1, 3, 4, 7, 8, 12, 14, 22, 32, 34, 46, 76, 189}

	t.Run("Test 1", func(t *testing.T) {
		got := search(76, testArray)
		want := 11

		assertInt(t, got, want)
	})
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
