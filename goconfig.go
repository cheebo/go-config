package go_config

import (
	"errors"
	"reflect"
)

type config struct {
	sources []Source
}

var (
	GoConfig Config

	NoVariablesInitialised = errors.New("no variables initialised")
)

func init() {
	GoConfig = New()
}

func New() Config {
	return &config{
		sources: []Source{},
	}
}

func (gc *config) UseSource(sources ...Source) {
	gc.sources = append(gc.sources, sources...)
}

func (gc *config) Unmarshal(v interface{}, prefix string) error {
	return NoVariablesInitialised
}

func (gc *config) Get(key string) interface{} {
	var value interface{}
	for _, src := range gc.sources {
		val := src.Get(key)
		// @todo check nil correct
		if val == nil {
			continue
		}
		value = val
	}
	return value
}

func (gc *config) Bool(key string) (bool, error) {
	val := gc.Get(key)
	if reflect.ValueOf(val).IsNil() {
		return false, NoVariablesInitialised
	}
	return val.(bool), nil
}

func (gc *config) Float(key string) (float64, error) {
	val := gc.Get(key)
	if reflect.ValueOf(val).IsNil() {
		return 0, NoVariablesInitialised
	}
	return val.(float64), nil
}

func (gc *config) Int(key string) (int, error) {
	val := gc.Get(key)
	if reflect.ValueOf(val).IsNil() {
		return 0, NoVariablesInitialised
	}
	return val.(int), nil
}

func (gc *config) UInt(key string) (uint, error) {
	val := gc.Get(key)
	if reflect.ValueOf(val).IsNil() {
		return 0, NoVariablesInitialised
	}
	return val.(uint), nil
}

func (gc *config) Slice(key, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (gc *config) String(key string) (string, error) {
	val := gc.Get(key)
	if reflect.ValueOf(val).IsNil() {
		return "", NoVariablesInitialised
	}
	return val.(string), nil
}

func (gc *config) StringMap(key string) map[string]interface{} {
	return map[string]interface{}{}
}

func (gc *config) IsSet(key string) bool {
	for _, src := range gc.sources {
		if src.IsSet(key) {
			return true
		}
	}
	return false
}
