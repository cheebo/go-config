package go_config

import (
	"reflect"
)

type Config interface {
	UseSource(sources ...Source)
	Unmarshal(v interface{}, prefix string) error

	Get(key string) interface{}

	Bool(key string) bool
	Float(key string) float64
	Int(key string) int
	UInt(key string) uint
	Slice(key, delimiter string) []interface{}
	String(key string) string
	StringMap(key string) map[string]interface{}

	IsSet(key string) bool
}

type Source interface {
	Get(key string) interface{}

	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	Int(key string) (int, error)
	UInt(key string) (uint, error)
	Slice(key, delimiter string) ([]interface{}, error)
	String(key string) (string, error)
	StringMap(key string) map[string]interface{}

	IsSet(key string) bool
}
