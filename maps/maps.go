package maps

import (
	"errors"
)

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
	ErrWordNotExist = DictionaryErr("cannot update word because it does not exist")
)

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

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case errors.Is(err, nil):
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordNotExist
	case errors.Is(err, nil):
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
