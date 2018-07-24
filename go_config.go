package go_config

import (
	"errors"
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
	// @todo implement
	return NoVariablesInitialised
}

func (gc *config) Get(key string) interface{} {
	var value interface{}
	for _, src := range gc.sources {
		val := src.Get(key)
		if val == nil {
			continue
		}
		value = val
	}
	return value
}

func (gc *config) Bool(key string) bool {
	val := gc.Get(key)
	if val == nil {
		return false
	}
	return val.(bool)
}

func (gc *config) Float(key string) float64 {
	val := gc.Get(key)
	if val == nil {
		return 0
	}
	return val.(float64)
}

func (gc *config) Int(key string) int {
	val := gc.Get(key)
	if val == nil {
		return 0
	}
	return val.(int)
}

func (gc *config) UInt(key string) uint {
	val := gc.Get(key)
	if val == nil {
		return 0
	}
	return val.(uint)
}

func (gc *config) Slice(key, delimiter string) []interface{} {
	val := gc.Get(key)
	if val == nil {
		return []interface{}{}
	}
	return val.([]interface{})
}

func (gc *config) String(key string) string {
	val := gc.Get(key)
	if val == nil {
		return ""
	}
	return val.(string)
}

func (gc *config) StringMap(key string) map[string]interface{} {
	val := gc.Get(key)
	if val == nil {
		return map[string]interface{}{}
	}
	return val.(map[string]interface{})
}

func (gc *config) IsSet(key string) bool {
	for _, src := range gc.sources {
		if src.IsSet(key) {
			return true
		}
	}
	return false
}
