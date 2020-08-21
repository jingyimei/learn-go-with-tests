package learn_go_with_tests

type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("key already exists, quit")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string, err error) {
	value, ok := d[key]
	if !ok {
		err = ErrNotFound
	}
	return value, err
}

func (d Dictionary) Add(key string, value string) (err error) {
	_, ok := d[key]
	if ok {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(word string, definition string) (err error){
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) (err error){
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return
}
