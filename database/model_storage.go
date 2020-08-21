package database

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type opts struct {
	model interface{}
	field string
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

// Get ...
func (mos *ModelStorage) Get() (interface{}, error) {
	mos.Lock()
	defer mos.Unlock()

	valueOfModel := reflect.ValueOf(mos.model)

	s := valueOfModel.Elem()

	if s.Kind() != reflect.Struct {
		return nil, errors.New("Should be structure model")
	}

	field := s.FieldByName(mos.field)

	for _, value := range mos.storage.Data {
		fmt.Println(reflect.TypeOf(value))
		fmt.Println(reflect.TypeOf(mos.Model))
		if reflect.TypeOf(value) == reflect.TypeOf(mos.Model) {
			localValue := reflect.ValueOf(value)

			sValue := localValue.Elem()

			fieldLocal := sValue.FieldByName(mos.field)

			if s.Kind() != reflect.Struct {
				return nil, errors.New("Should be structure model")
			}
			switch field.Kind() {
			case reflect.Int:
				fmt.Println(0)
				valueCompareLocal := int(fieldLocal.Int())

				if valueCompareLocal == int(field.Int()) {
					return value, nil
				}
			case reflect.Uint:
				fmt.Println(1)
				valueCompareLocal := uint(fieldLocal.Uint())

				if valueCompareLocal == uint(field.Uint()) {
					return value, nil
				}

			case reflect.String:
				valueCompareLocal := fieldLocal.String()

				if valueCompareLocal == field.String() {
					return value, nil
				}
			}

		}
	}

	return nil, errors.New("Value not found")

}
