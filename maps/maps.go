package maps

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("definition already exists!")
	ErrWordDoesNotExists = DictionaryErr("definition not updated, word not found")
	ErrDelete            = DictionaryErr("word not deleted; it doesn't exist already!")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrDelete
	default:
		return err
	}
	return nil
}
