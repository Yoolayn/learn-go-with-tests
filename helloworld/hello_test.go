package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Magda", "")
	want := "Hello, Magda!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Magda", "")
		want := "Hello, Magda!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Beatrice", "French")
		want := "Bonjour, Beatrice!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
