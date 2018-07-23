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

func (gc *config) Bool(key string) bool {
	val := gc.Get(key)
	if !reflect.ValueOf(val).CanAddr() {
		return false
	}
	return val.(bool)
}

func (gc *config) Float(key string) float64 {
	val := gc.Get(key)
	if !reflect.ValueOf(val).CanAddr() {
		return 0
	}
	return val.(float64)
}

func (gc *config) Int(key string) int {
	val := gc.Get(key)
	if !reflect.ValueOf(val).CanAddr() {
		return 0
	}
	return val.(int)
}

func (gc *config) UInt(key string) uint {
	val := gc.Get(key)
	if !reflect.ValueOf(val).CanAddr() {
		return 0
	}
	return val.(uint)
}

func (gc *config) Slice(key, delimiter string, kind reflect.Kind) []interface{} {
	return []interface{}{}
}

func (gc *config) String(key string) string {
	val := gc.Get(key)
	if !reflect.ValueOf(val).CanAddr() {
		return ""
	}
	return val.(string)
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
