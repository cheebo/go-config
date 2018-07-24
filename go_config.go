package go_config

import (
	"errors"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type config struct {
	sources []Source
}

var (
	GoConfig Config

	NoVariablesInitialised = errors.New("no variables initialised")
	NotAStructPtr          = errors.New("expects pointer to a struct")
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
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return NotAStructPtr
	}
	ref := ptr.Elem()
	if ref.Kind() != reflect.Struct {
		return NotAStructPtr
	}

	return gc.setup(v, prefix)
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

func (gc *config) setup(v interface{}, parent string) error {
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
			gc.setup(refField.Addr().Interface(), name)
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
			// @todo fill slice
		case reflect.Map:
			// @todo fill map
		}

	}
	return nil
}
