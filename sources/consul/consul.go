package consul

import (
	"bytes"
	"fmt"
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type Consul struct {
	Path      string
	Type      go_config.ConfigType
	Namespace string
}

type consul struct {
	prefix string
	config types.ConsulConfig
	data   map[string]interface{}
}

func Source(log *logrus.Logger, config types.ConsulConfig, dataSource ...Consul) (go_config.Source, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	for _, ds := range dataSource {
		q := &api.QueryOptions{
			Datacenter:        config.Datacenter,
			Token:             config.Token,
			RequireConsistent: true,
		}
		kvpair, _, err := client.KV().Get(ds.Path, q)
		if err != nil {
			log.WithField("component", "go-config.consul.Source").Error(fmt.Sprintf("Can't read config from %s err: %s", ds.Path, err.Error()))
			continue
		}
		if kvpair == nil {
			log.WithField("component", "go-config.consul.Source").Error(fmt.Sprintf("Can't read config from %s", ds.Path))
			continue
		}

		m := map[string]interface{}{}
		err = go_config.ReadConfig(bytes.NewBuffer([]byte(strings.TrimSpace(string(kvpair.Value)))), ds.Type, m)
		if err != nil {
			log.WithField("component", "go-config.consul.Source").Error(fmt.Sprintf("Can't parse config from %s type %s", ds.Path, ds.Type))
			continue
		}

		err = go_config.MergeMapWithPath(data, m, strings.Split(ds.Namespace, "."))
		if err != nil {
			log.Error(err)
		}
	}

	return &consul{
		config: config,
		data:   data,
	}, nil
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

func (c *consul) Slice(key, delimiter string) ([]interface{}, error) {
	val := c.lookup(c.data, c.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(val)
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
