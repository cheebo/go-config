package env

import (
	"github.com/cheebo/go-config"
	"github.com/spf13/cast"
	"os"
	"strconv"
	"strings"
)

type env struct {
	prefix string
}

func Source(prefix string) go_config.Source {
	return &env{
		prefix: prefix,
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
	val, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func (e *env) Slice(key, delimiter string) ([]interface{}, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(strings.Split(val, delimiter))
}

func (e *env) String(key string) (string, error) {
	val, ok := os.LookupEnv(e.key(key))
	if !ok {
		return "", go_config.NoVariablesInitialised
	}
	return val, nil
}

func (e *env) StringMap(key string) map[string]interface{} {
	// @todo implement
	return map[string]interface{}{}
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
