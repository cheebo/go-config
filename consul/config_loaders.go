package consul

import (
	"errors"
	"github.com/hashicorp/consul/api"
	"github.com/cheebo/consul-utils/kv"
	consultypes "github.com/cheebo/consul-utils/types"
	"github.com/cheebo/go-config/types"
)

func GetDatabaseConfig(client *api.Client, path string, opt consultypes.QueryOptions) (*types.Database, error) {
	var cfg types.Database
	str, err := kv.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return cfg, nil
}

func GetRedisConfig(client *api.Client, path string, opt consultypes.QueryOptions) (*types.Redis, error) {
	var cfg types.Redis
	str, err := kv.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return cfg, nil
}

func GetOAuthProvidersConfig(client *api.Client, path string, opt consultypes.QueryOptions) ([]types.OAuthProvider, error) {
	var cfg []types.OAuthProvider
	str, err := kv.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return cfg, nil
}