package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected an error")
		}

		assertStrings(t, err.Error(), want.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		_ = d.Add("test", "this is just a test")

		want := "this is just a test"

		assertDefinition(t, d, "test", want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefintion := "new definition"

		_ = dictionary.Update(word, newDefintion)

		assertDefinition(t, dictionary, word, newDefintion)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefintion := "new definition"

		err := dictionary.Update(word, newDefintion)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		_ = dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("expected %q to be deleted", word)
		}
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrDelete)
	})
}
