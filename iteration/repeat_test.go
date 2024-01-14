package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
	t.Run("more", func(t *testing.T) {
		repeated := Repeat("a", 15)
		expected := "aaaaaaaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	b.Run("five", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 5)
		}
	})
	b.Run("fifteen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 15)
		}
	})
}

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}
