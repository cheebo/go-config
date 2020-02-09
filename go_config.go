package go_config

import (
	"errors"
	"reflect"
	"strings"

	"github.com/cheebo/go-config/internal/utils"
	utils2 "github.com/cheebo/go-config/pkg/utils"
	"github.com/spf13/cast"
)

type config struct {
	sub      string
	sources  []Source
	defaults map[string]interface{}
	config   map[string]interface{}
}

var (
	NoVariablesInitialised = errors.New("no variables initialised")
	NotAStructPtr          = errors.New("expects pointer to a struct")
	NotAStruct             = errors.New("expects a struct")
)

func New() Config {
	return newConfig("", []Source{}, map[string]interface{}{})
}

func newConfig(sub string, sources []Source, defaults map[string]interface{}) Config {
	return &config{
		sub:      sub,
		sources:  sources,
		defaults: defaults,
	}
}

func (gc *config) UseSource(sources ...Source) {
	gc.sources = append(gc.sources, sources...)
}

func (gc *config) Unmarshal(v interface{}, prefix string) error {
	if t := reflect.TypeOf(v); t.Kind() != reflect.Ptr {
		return NotAStructPtr
	}

	refVal := reflect.ValueOf(v)
	for refVal.Kind() == reflect.Ptr && !refVal.IsNil() {
		refVal = refVal.Elem()
	}

	if refVal.Kind() != reflect.Struct {
		return NotAStruct
	}

	refType := reflect.TypeOf(v)
	for refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
	}

	for i := 0; i < refVal.NumField(); i++ {
		field := refType.Field(i)
		refField := refVal.Field(i)

		name := strings.ToLower(field.Name)
		if len(prefix) > 0 {
			name = prefix + "." + name
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
			gc.Unmarshal(refField.Addr().Interface(), name)
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
			switch refField.Type().Elem().Kind() {
			case reflect.Int:
				refField.Set(reflect.ValueOf(cast.ToIntSlice(gc.Get(name))))
			case reflect.String:
				refField.Set(reflect.ValueOf(cast.ToStringSlice(gc.Get(name))))
			case reflect.Bool:
				refField.Set(reflect.ValueOf(cast.ToBoolSlice(gc.Get(name))))
			case reflect.Interface:
				refField.Set(reflect.ValueOf(cast.ToSlice(gc.Get(name))))
			}
		case reflect.Map:
			switch refField.Type().Elem().Kind() {
			case reflect.Bool:
				refField.Set(reflect.ValueOf(cast.ToStringMapBool(gc.Get(name))))
			case reflect.Int:
				refField.Set(reflect.ValueOf(cast.ToStringMapInt(gc.Get(name))))
			case reflect.String:
				refField.Set(reflect.ValueOf(cast.ToStringMapString(gc.Get(name))))
			case reflect.Interface:
				refField.Set(reflect.ValueOf(cast.ToStringMap(gc.Get(name))))
			case reflect.Slice:
				if refField.Type().Elem().Elem().Kind() == reflect.String {
					refField.Set(reflect.ValueOf(cast.ToStringMapStringSlice(gc.Get(name))))
				}
			}
		}

	}
	return nil
}

func (gc *config) SetDefault(key string, val interface{}) {
	path := strings.Split(key, ".")
	m := map[string]interface{}{}

	if t := reflect.TypeOf(val); t.Kind() == reflect.Struct {
		val = utils2.StructToStringMap(val)
	}

	m[path[len(path)-1]] = val
	utils.MergeMapWithPath(gc.defaults, m, path[:len(path)-1])
}

func (gc *config) Get(key string) interface{} {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val := src.Get(key)
		if val == nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
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
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Float(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return value.(float64)
}

func (gc *config) Int(key string) int {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Int(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToInt(value)
}

func (gc *config) Int8(key string) int8 {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Int8(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToInt8(value)
}

func (gc *config) Int32(key string) int32 {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Int32(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToInt32(value)
}

func (gc *config) Int64(key string) int64 {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Int64(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToInt64(value)
}

func (gc *config) UInt(key string) uint {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.UInt(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return value.(uint)
}

func (gc *config) UInt32(key string) uint32 {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.UInt32(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToUint32(value)
}

func (gc *config) UInt64(key string) uint64 {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.UInt64(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
		if value == nil {
			return 0
		}
	}
	return cast.ToUint64(value)
}

func (gc *config) Sub(key string) Fields {
	key = nestedKey(gc.sub, key)
	c := newConfig(key, gc.sources, gc.defaults)
	return c
}

func (gc *config) Slice(key string) []interface{} {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.Slice(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
	}
	return cast.ToSlice(value)
}

func (gc *config) SliceInt(key string) []int {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.SliceInt(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
	}
	return cast.ToIntSlice(value)
}

func (gc *config) SliceString(key string) []string {
	key = nestedKey(gc.sub, key)
	var value interface{}
	for _, src := range gc.sources {
		val, err := src.SliceString(key)
		if err != nil {
			continue
		}
		value = val
	}
	if value == nil {
		value = utils.Lookup(gc.defaults, strings.Split(key, "."))
	}
	return cast.ToStringSlice(value)
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
	return cast.ToStringMap(val)
}

func (gc *config) StringMapInt(key string) map[string]int {
	val := gc.Get(key)
	if val == nil {
		return map[string]int{}
	}
	return cast.ToStringMapInt(val)
}

func (gc *config) StringMapSliceString(key string) map[string][]string {
	val := gc.Get(key)
	if val == nil {
		return map[string][]string{}
	}
	return cast.ToStringMapStringSlice(val)
}

func (gc *config) StringMapString(key string) map[string]string {
	val := gc.Get(key)
	if val == nil {
		return map[string]string{}
	}
	return cast.ToStringMapString(val)
}

func (gc *config) IsSet(key string) bool {
	key = nestedKey(gc.sub, key)
	for _, src := range gc.sources {
		if src.IsSet(key) {
			return true
		}
	}
	return false
}

func nestedKey(sub, key string) string {
	if len(sub) > 0 {
		return sub + "." + key
	}
	return key
}
