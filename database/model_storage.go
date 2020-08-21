package database

import (
	"errors"
	"reflect"
	"sync"
)

type opts struct {
	model interface{}
	field string
	condition
}

type condition struct {
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
func (mos *ModelStorage) Get() (interface{}, error) {
	mos.Lock()
	defer mos.Unlock()

	for _, value := range mos.storage.Data {
		if reflect.TypeOf(value) == reflect.TypeOf(mos.model) {
			localValue := reflect.ValueOf(value)

			sValue := localValue.Elem()

			fieldLocal := sValue.FieldByName(mos.field)

			if sValue.Kind() != reflect.Struct {
				return nil, errors.New("Should be structure model")
			}

			switch fieldLocal.Kind() {
			case reflect.Int:
				valueCompareLocal := int(fieldLocal.Int())

				if valueCompareLocal == mos.opts.equal.(int) {
					return value, nil
				}
			case reflect.Uint:
				valueCompareLocal := uint(fieldLocal.Uint())

				if reflect.TypeOf(mos.opts.equal).Kind() == reflect.Int {
					valueCompare := mos.opts.equal.(int)

					if valueCompareLocal == uint(valueCompare) {
						return value, nil
					}
					continue
				}

				if valueCompareLocal == mos.opts.equal.(uint) {
					return value, nil
				}

			case reflect.String:
				valueCompareLocal := fieldLocal.String()

				if valueCompareLocal == mos.opts.equal.(string) {
					return value, nil
				}
			}

		}
	}

	return nil, errors.New("Value not found")

}

// Set ...
func (mos *ModelStorage) Set() {
	mos.storage.Write(mos.model)
}

// SetArray ...
func (mos *ModelStorage) SetArray(data []interface{}) {
	mos.Lock()
	defer mos.Unlock()

	mos.storage.WriteArray(data)
}
