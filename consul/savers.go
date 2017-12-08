package consul

import (
	"encoding/json"

	"github.com/cheebo/consul-utils"
	"github.com/hashicorp/consul/api"
)

func SetConfig(client *api.Client, path string, opt consul_utils.QueryOptions, config interface{}) error {
	var err error
	cfg, ok := config.(string)
	if !ok {
		cfgBytes, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			return ErrCantMarshalJSON
		}
		cfg = string(cfgBytes)
	}

	_, err = consul_utils.PutKV(client, path[1:], cfg, opt)
	if err != nil {
		return ErrCantPutConfig
	}
	return nil
}
