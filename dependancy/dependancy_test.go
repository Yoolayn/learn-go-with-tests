package dependancy

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Magda")

		got := buffer.String()
		want := "Hello, Magda"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
