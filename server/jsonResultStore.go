package main

import (
	"os"
	"encoding/json"
)

type jsonResultStore struct {
	storagePath string
}

func (c *jsonResultStore) getAll() ([]payload, error) {
	file, err := os.OpenFile(c.storagePath, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var result []payload
	json.NewDecoder(file).Decode(&result)

	return result, nil
}

func NewResultStore(storagePath string) *jsonResultStore {
	return &jsonResultStore{storagePath:storagePath}
}

func (c *jsonResultStore) save(p payload) error{
	file, err := os.OpenFile(c.storagePath, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	var data []payload
	json.NewDecoder(file).Decode(&data)
	data = append(data, p)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	file.WriteAt(bytes, 0)
	return nil
}

