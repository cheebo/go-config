package go_config

import (
	"errors"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type config struct {
	sub string
	sources  []Source
	defaults map[string]interface{}
	config map[string]interface{}
}

var (
	NoVariablesInitialised = errors.New("no variables initialised")
	NotAStructPtr          = errors.New("expects pointer to a struct")
)

func New() Config {
	return newConfig("", []Source{}, map[string]interface{}{})
}

func newConfig(sub string, sources []Source, defaults map[string]interface{}) Config {
	return &config{
		sub: sub,
		sources:  sources,
		defaults: defaults,
	}
}

func (gc *config) UseSource(sources ...Source) {
	gc.sources = append(gc.sources, sources...)
}

func (gc *config) Unmarshal(v interface{}, prefix string) error {
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return NotAStructPtr
	}
	ref := ptr.Elem()
	if ref.Kind() != reflect.Struct {
		return NotAStructPtr
	}

	return gc.unmarshal(v, prefix)
}

func (gc *config) SetDefault(key string, val interface{}) {
	path := strings.Split(key, ".")
	m := map[string]interface{}{}
	m[path[len(path)-1]] = val
	MergeMapWithPath(gc.defaults, m, path[:len(path)-1])
}

func (gc *config) Get(key string) interface{} {
	if len(gc.sub) > 0 {
		key = gc.sub + "." + key
	}
	var value interface{}
	for _, src := range gc.sources {
		val := src.Get(key)
		if val == nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = Lookup(gc.defaults, strings.Split(key, "."))
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
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Float(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return value.(float64)
}

func (gc *config) Int(key string) int {
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Int(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return value.(int)
}

func (gc *config) UInt(key string) uint {
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.UInt(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return value.(uint)
}

func (gc *config) Sub(key string) Fields {
	if len(gc.sub) > 0 {
		key = gc.sub + "." + key
	}
	c := newConfig(key, gc.sources, gc.defaults)
	return c
}

func (gc *config) Slice(key, delimiter string) []interface{} {
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Slice(key, delimiter)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = Lookup(gc.defaults, strings.Split(key, "."))
	}
	return cast.ToSlice(value)
}

func (gc *config) String(key string) string {
	val := gc.Get(key)
	if val == nil {
		return ""
	}
	v, ok := val.(string)
	if !ok {
		return ""
	}
	return v
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

func (gc *config) unmarshal(v interface{}, parent string) error {
	refVal := reflect.ValueOf(v)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	if refVal.Kind() != reflect.Struct {
		return nil
	}

	refType := reflect.TypeOf(refVal.Interface())

	for i := 0; i < refVal.NumField(); i++ {
		field := refType.Field(i)
		refField := refVal.Field(i)

		name := strings.ToLower(field.Name)
		if len(parent) > 0 {
			name = parent + "." + name
		}

		if refField.Kind() == reflect.Ptr {
			if refField.IsNil() {
				refField = reflect.New(refField.Type().Elem())
				refVal.Field(i).Set(refField)
				refField = refField.Elem()
			} else {
				refField = refField.Elem()
			}
		}

		if refField.Kind() == reflect.Struct {
			gc.unmarshal(refField.Addr().Interface(), name)
			continue
		}

		if !refField.CanSet() {
			continue
		}

		switch refField.Kind() {
		case reflect.Int:
			refField.SetInt(cast.ToInt64(gc.Get(name)))
		case reflect.Uint:
			refField.SetUint(cast.ToUint64(gc.Get(name)))
		case reflect.Float64:
			refField.SetFloat(cast.ToFloat64(gc.Get(name)))
		case reflect.Bool:
			refField.SetBool(cast.ToBool(gc.Get(name)))
		case reflect.String:
			refField.SetString(cast.ToString(gc.Get(name)))
		case reflect.Slice:
			m := cast.ToSlice(gc.Get(name))
			refField.Set(reflect.ValueOf(m))
		case reflect.Map:
			m := cast.ToStringMap(gc.Get(name))
			refField.Set(reflect.ValueOf(m))
		}

	}
	return nil
}
