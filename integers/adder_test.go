package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assetCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	sum := Add(2, 2)
	expected := 4

	assetCorrectMessage(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
