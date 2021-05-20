package models

import "errors"

var Storage StorageModel

//StorageModel ...
type StorageModel map[string]string

func (m StorageModel) Create(key string, value string) error {
	Storage[key] = value
	return nil
}

//Get ...
func (m StorageModel) Get(key string) (string, error) {
	pair, found := Storage[key]
	if !found {
		return "", errors.New("key not found")
	}
	return pair, nil
}

//Delete ...
func (m StorageModel) Delete(key string) error {
	_, found := Storage[key]
	if !found {
		return errors.New("key not found")
	}

	delete(Storage, key)

	return nil
}
