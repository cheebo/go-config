package flag

import (
	"flag"
	"github.com/cheebo/go-config"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type flags struct {
	fs     *flag.FlagSet
	values map[string]interface{}
}

func Source() go_config.Source {
	return &flags{
		fs:     flag.NewFlagSet("", flag.ContinueOnError),
		values: make(map[string]interface{}),
	}
}

func (f *flags) Get(key string) interface{} {
	return f.lookup(f.key(key))
}

func (f *flags) Bool(key string) (bool, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (f *flags) Float(key string) (float64, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToFloat64E(val)
}

func (f *flags) Int(key string) (int, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToIntE(val)
}

func (f *flags) UInt(key string) (uint, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUintE(val)
}

func (f *flags) Slice(key, delimiter string) ([]interface{}, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(strings.Split(val.(string), delimiter))
}

func (f *flags) String(key string) (string, error) {
	val := f.lookup(f.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (f *flags) StringMap(key string) map[string]interface{} {
	val := f.lookup(f.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (f *flags) IsSet(key string) bool {
	val := f.lookup(f.key(key))
	if val == nil {
		return false
	}
	return true
}

func (f *flags) key(key string) string {
	return strings.ToLower(key)
}

func (f *flags) lookup(flag string) interface{} {
	// @todo implement method: get flag value
	return nil
}
