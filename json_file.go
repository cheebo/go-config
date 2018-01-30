package go_config

import (
	"github.com/cheebo/go-config/utils"
	"io/ioutil"
	"reflect"
	"strings"
)

type jsonFIle struct {
	path string
	data map[string]interface{}
}

func JsonFileSource(path string) Source {
	return &jsonFIle{
		path: path,
		data: map[string]interface{}{},
	}
}

func (self *jsonFIle) Init(vals map[string]*Variable) error {
	data, err := ioutil.ReadFile(self.path)
	if err != nil {
		return err
	}
	self.data, err = utils.JsonParse(data)
	return err
}

func (self *jsonFIle) Int(name string) (int, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case int:
		return val.(int), nil
	case float64:
		return int(val.(float64)), nil
	}
	return 0, nil
}

func (self *jsonFIle) Float(name string) (float64, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case float64:
		return val.(float64), nil
	}
	return 0, nil
}

func (self *jsonFIle) UInt(name string) (uint, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	switch val.(type) {
	case uint:
		return val.(uint), nil
	case float64:
		return uint(val.(float64)), nil
	}
	return 0, nil
}

func (self *jsonFIle) String(name string) (string, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return "", nil
	}
	switch val.(type) {
	case string:
		return val.(string), nil
	}
	return "", nil
}

func (self *jsonFIle) Bool(name string) (bool, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return false, nil
	}
	switch val.(type) {
	case bool:
		return val.(bool), nil
	}
	return false, nil
}

func (self *jsonFIle) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (self *jsonFIle) name(name string) string {
	return strings.ToLower(name)

}
