package consul

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type DataSource struct {
	Path      string
	Type      go_config.ConfigType
	Namespace string
}

type consul struct {
	prefix string
	config types.ConsulConfig
	data   map[string]interface{}
}

func Source(config types.ConsulConfig, dataSource ...DataSource) go_config.Source {
	for _, ds := range dataSource {
		// @todo import data from consul into local map
		if len(ds.Namespace) == 0 {
			continue
		}
		// @todo get data, ReadConfig and save to map
	}
	return &consul{
		config: config,
		data:   make(map[string]interface{}),
	}
}

func (c *consul) Get(key string) interface{} {
	return c.lookup(c.data, c.key(key))
}

func (c *consul) Bool(key string) (bool, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (c *consul) Float(key string) (float64, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToFloat64E(val)
}

func (c *consul) Int(key string) (int, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToIntE(val)
}

func (c *consul) UInt(key string) (uint, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUintE(val)
}

func (c *consul) Slice(key, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (c *consul) String(key string) (string, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (c *consul) StringMap(key string) map[string]interface{} {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (c *consul) IsSet(key string) bool {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return false
	}
	return true
}

func (c *consul) key(key string) []string {
	return strings.Split(key, ".")
}

func (c *consul) lookup(source map[string]interface{}, key []string) interface{} {
	if len(key) == 0 {
		return source
	}

	next, ok := source[key[0]]
	if ok {
		if len(key) == 1 {
			return next
		}

		// Nested case
		switch next.(type) {
		case map[interface{}]interface{}:
			return c.lookup(cast.ToStringMap(next), key[1:])
		case map[string]interface{}:
			return c.lookup(next.(map[string]interface{}), key[1:])
		default:
			return nil
		}
	}
	return nil
}
