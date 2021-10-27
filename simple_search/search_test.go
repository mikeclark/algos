package simplesearch

import "testing"

func TestSearch(t *testing.T) {
	// t.Fatal("not implemented")
	testArray := []int{9, 1, 33, 21, 78, 42, 4}

	t.Run("Test 1", func(t *testing.T) {
		got := search(78, testArray)
		want := 4

		assertInt(t, got, want)
	})
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
