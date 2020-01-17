package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/cheebo/go-config"
	"github.com/spf13/cast"
)

type env struct {
	prefix    string
	delimiter string
}

func Source(prefix string, delimiter string) go_config.Source {
	return &env{
		prefix:    prefix,
		delimiter: delimiter,
	}
}

func (e *env) Get(key string) interface{} {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return nil
	}
	return val
}

func (e *env) Bool(key string) (bool, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return false, go_config.NoVariablesInitialised
	}
	return strconv.ParseBool(val)
}

func (e *env) Int(key string) (int, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return strconv.Atoi(val)
}

func (e *env) Int8(key string) (int8, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt8E(val)
}

func (e *env) Int32(key string) (int32, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt32E(val)
}

func (e *env) Int64(key string) (int64, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt64E(val)
}

func (e *env) Float(key string) (float64, error) {
	v, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (e *env) UInt(key string) (uint, error) {
	v, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUintE(v)
}

func (e *env) UInt32(key string) (uint32, error) {
	v, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUint32E(v)
}

func (e *env) UInt64(key string) (uint64, error) {
	v, ok := os.LookupEnv(e.key(key))
	if !ok {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUint64E(v)
}

func (e *env) Slice(key string) ([]interface{}, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	var slice []interface{}
	for _, s := range strings.Split(val, e.delimiter) {
		slice = append(slice, s)
	}
	return slice, nil
}

func (e *env) SliceInt(key string) ([]int, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return []int{}, go_config.NoVariablesInitialised
	}
	var slice []string
	for _, s := range strings.Split(val, e.delimiter) {
		slice = append(slice, s)
	}
	return cast.ToIntSliceE(slice)
}

func (e *env) SliceString(key string) ([]string, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return []string{}, go_config.NoVariablesInitialised
	}
	var slice []interface{}
	for _, s := range strings.Split(val, e.delimiter) {
		slice = append(slice, s)
	}
	return cast.ToStringSliceE(slice)
}

func (e *env) String(key string) (string, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return "", go_config.NoVariablesInitialised
	}
	return val, nil
}

func (e *env) StringMap(key string) map[string]interface{} {
	return map[string]interface{}{}
}

func (e *env) StringMapInt(key string) map[string]int {
	return map[string]int{}
}

func (e *env) StringMapSliceString(key string) map[string][]string {
	return map[string][]string{}
}

func (e *env) StringMapString(key string) map[string]string {
	return map[string]string{}
}

func (e *env) IsSet(key string) bool {
	_, ok := os.LookupEnv(e.key(key))
	return ok
}

func (e *env) key(key string) string {
	key = strings.Replace(strings.ToUpper(key), ".", "_", -1)
	if len(e.prefix) > 0 {
		return strings.ToUpper(e.prefix) + "_" + key
	}
	return key
}
