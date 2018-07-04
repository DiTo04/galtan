package data

import (
	"os"
	"encoding/json"
)

type jsonResultStore struct {
	storagePath string
}

func (c *jsonResultStore) GetAll() ([]Payload, error) {
	file, err := os.OpenFile(c.storagePath, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var result []Payload
	json.NewDecoder(file).Decode(&result)

	return result, nil
}

func NewResultStore(storagePath string) *jsonResultStore {
	return &jsonResultStore{storagePath:storagePath}
}

func (c *jsonResultStore) Save(p Payload) error{
	file, err := os.OpenFile(c.storagePath, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	var data []Payload
	json.NewDecoder(file).Decode(&data)
	data = append(data, p)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	file.WriteAt(bytes, 0)
	return nil
}

