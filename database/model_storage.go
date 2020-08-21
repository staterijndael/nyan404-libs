package database

import (
	"errors"
	"reflect"
	"sync"
)

type opts struct {
	model interface{}
	field interface{}
	equal interface{}
}

// ModelStorage ...
type ModelStorage struct {
	storage *MemoryStorage
	opts
	sync.Mutex
}

// NewModelStorage ...
func NewModelStorage() *ModelStorage {
	return &ModelStorage{
		storage: NewMemoryStorage(),
	}
}

// Model ...
func (mos *ModelStorage) Model(model interface{}) *ModelStorage {
	mos.Lock()
	mos.model = model
	mos.Unlock()

	return mos
}

// Field ...
func (mos *ModelStorage) Field(field string) *ModelStorage {
	mos.Lock()
	mos.field = field
	mos.Unlock()

	return mos
}

// Equal ...
func (mos *ModelStorage) Equal(value interface{}) *ModelStorage {
	mos.Lock()
	mos.equal = value
	mos.Unlock()

	return mos
}

// Get ...
func (mos *ModelStorage) Get() error {
	mos.Lock()
	defer mos.Unlock()

	typeOfModel := reflect.TypeOf(mos.model)
	s := typeOfModel.Elem()

	if s.Kind() != reflect.Struct {
		return errors.New("Should be structure model")
	}

	field := s.FieldByName(mos.Field)

}
