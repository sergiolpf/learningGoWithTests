package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assetCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q but got %q", got, want)
		}
	}


	repeated := Repeat("a")
	expected := "aaaaa"

	assetCorrectMessage(t, repeated, expected)

}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i< b.N; i++{
		Repeat("a")

	}
}


