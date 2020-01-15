package consul

import (
	"bytes"
	"strings"

	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/internal/reader"
	"github.com/cheebo/go-config/internal/utils"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cast"
)

type Consul struct {
	Path      string
	Type      reader.ConfigType
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
		err = reader.ReadConfig(bytes.NewBuffer([]byte(strings.TrimSpace(string(kvpair.Value)))), ds.Type, m)
		if err != nil {
			// todo: return or log error
			continue
		}

		err = utils.MergeMapWithPath(data, m, strings.Split(ds.Namespace, "."))
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
	return utils.Lookup(c.data, c.key(key))
}

func (c *consul) Bool(key string) (bool, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (c *consul) Float(key string) (float64, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToFloat64E(val)
}

func (c *consul) Int(key string) (int, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToIntE(val)
}

func (c *consul) Int8(key string) (int8, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt8E(val)
}

func (c *consul) Int32(key string) (int32, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt32E(val)
}

func (c *consul) Int64(key string) (int64, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToInt64E(val)
}

func (c *consul) UInt(key string) (uint, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUintE(val)
}

func (c *consul) UInt32(key string) (uint32, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUint32E(val)
}

func (c *consul) UInt64(key string) (uint64, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}
	return cast.ToUint64E(val)
}

func (c *consul) Slice(key string) ([]interface{}, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(val)
}

func (c *consul) SliceInt(key string) ([]int, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return []int{}, go_config.NoVariablesInitialised
	}
	return cast.ToIntSliceE(val)
}

func (c *consul) SliceString(key string) ([]string, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return []string{}, go_config.NoVariablesInitialised
	}
	return cast.ToStringSliceE(val)
}

func (c *consul) String(key string) (string, error) {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (c *consul) StringMap(key string) map[string]interface{} {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (c *consul) StringMapInt(key string) map[string]int {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return map[string]int{}
	}
	return cast.ToStringMapInt(val)
}

func (c *consul) StringMapSliceString(key string) map[string][]string {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return map[string][]string{}
	}
	return cast.ToStringMapStringSlice(val)
}

func (c *consul) StringMapString(key string) map[string]string {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return map[string]string{}
	}
	return cast.ToStringMapString(val)
}

func (c *consul) IsSet(key string) bool {
	val := utils.Lookup(c.data, c.key(key))
	if val == nil {
		return false
	}
	return true
}

func (c *consul) key(key string) []string {
	return strings.Split(key, ".")
}
