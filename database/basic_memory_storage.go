package database

import (
	"errors"
	"sync"
)

// MemoryStorage ...
type MemoryStorage struct {
	Name string
	Data map[uint]interface{}
	sync.Mutex
}

// NewMemoryStorage ...
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Name: "",
		Data: make(map[uint]interface{}),
	}
}

// Write ...
func (mem *MemoryStorage) Write(data interface{}) {

	mem.Data[uint(len(mem.Data))] = data
}

// WriteArray ...
func (mem *MemoryStorage) WriteArray(data []interface{}) {

	for _, value := range data {
		mem.Write(value)
	}
}

// Read ...
func (mem *MemoryStorage) Read(id uint) (interface{}, error) {

	data := mem.Data[id]
	if data == nil {
		return nil, errors.New("Value doest not exist")
	}

	return data, nil
}
