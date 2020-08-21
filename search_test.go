package learn_go_with_tests

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("key exists", func (t *testing.T){
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("key does not exist", func (t *testing.T){
		dictionary := Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test1")
		assertStrings(t, got, "")
		assertError(t, err, ErrNotFound)
	})
}

func TestAddKey(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
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

func TestUpdateKey(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDeleteKey(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected '%s' to be deleted", word)
		}
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete("test1")
		if err != ErrWordDoesNotExist {
			t.Errorf("delete should throw an error, because test1 does not exist")
		}

		_, err = dictionary.Search(word)
		if err != nil {
			t.Errorf("Expected '%s' not to be deleted", word)
		}
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, err error, expectedErr error) {
	t.Helper()

	if err != expectedErr {
		t.Errorf("expect to get an error %q but got %q", expectedErr.Error(), err.Error())
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}
