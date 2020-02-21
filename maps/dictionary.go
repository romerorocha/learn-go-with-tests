package maps

type Dictionary map[string]string

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists in dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
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

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = definition
		return nil
	}

	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
