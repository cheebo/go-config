package go_config

import (
	"reflect"
)

type Config interface {
	Source

	UseSource(sources ...Source)
	Unmarshal(v interface{}, prefix string) error
}

type Source interface {
	Get(key string) interface{}

	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	Int(key string) (int, error)
	UInt(key string) (uint, error)
	Slice(key, delimiter string, kind reflect.Kind) ([]interface{}, error)
	String(key string) (string, error)
	StringMap(key string) map[string]interface{}

	IsSet(key string) bool
}
