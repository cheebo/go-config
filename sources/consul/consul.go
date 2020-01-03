package consul

import (
	"bytes"
	"github.com/cheebo/go-config"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cast"
	"strings"
)

type Consul struct {
	Path      string
	Type      go_config.ConfigType
	Namespace string
}

type consul struct {
	prefix string
	config ConsulConfig
	data   map[string]interface{}
}

func Source(config ConsulConfig, dataSource ...Consul) (go_config.Source, error) {
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
			// todo: return or log error
			continue
		}
		if kvpair == nil {
			// todo: return or log error
			continue
		}

		m := map[string]interface{}{}
		err = go_config.ReadConfig(bytes.NewBuffer([]byte(strings.TrimSpace(string(kvpair.Value)))), ds.Type, m)
		if err != nil {
			// todo: return or log error
			continue
		}

		err = go_config.MergeMapWithPath(data, m, strings.Split(ds.Namespace, "."))
		if err != nil {
			// todo: return or log error
		}
	}

	return &consul{
		config: config,
		data:   data,
	}, nil
}

func (c *consul) Get(key string) interface{} {
	return go_config.Lookup(c.data, c.key(key))
}

func (c *consul) Bool(key string) (bool, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (c *consul) Float(key string) (float64, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToFloat64E(val)
}

func (c *consul) Int(key string) (int, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToIntE(val)
}

func (c *consul) UInt(key string) (uint, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUintE(val)
}

func (c *consul) Slice(key, delimiter string) ([]interface{}, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(val)
}

func (c *consul) String(key string) (string, error) {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (c *consul) StringMap(key string) map[string]interface{} {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (c *consul) IsSet(key string) bool {
	val := go_config.Lookup(c.data, c.key(key))
	if val == nil {
		return false
	}
	return true
}

func (c *consul) key(key string) []string {
	return strings.Split(key, ".")
}
