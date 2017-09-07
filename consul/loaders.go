package consul

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	consulutils "github.com/cheebo/consul-utils"
	"github.com/cheebo/go-config/types"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GetAMQPConfig(client *api.Client, path string, opt consulutils.QueryOptions) (*types.AMQPConfig, error) {
	var cfg types.AMQPConfig
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return &cfg, nil
}

func GetDatabaseConfig(client *api.Client, path string, opt consulutils.QueryOptions) (*types.DatabaseService, error) {
	var cfg types.DatabaseService
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return &cfg, nil
}

func GetRedisConfig(client *api.Client, path string, opt consulutils.QueryOptions) (*types.RedisService, error) {
	var cfg types.RedisService
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return &cfg, nil
}

func GetOAuthProviderConfig(client *api.Client, path string, opt consulutils.QueryOptions) (types.OAuthProvider, error) {
	var cfg types.OAuthProvider
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return cfg, nil
}

func GetOAuthProviderListConfig(client *api.Client, path string, opt consulutils.QueryOptions) ([]types.OAuthProvider, error) {
	var cfg []types.OAuthProvider
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return cfg, nil
}

func GetRsaPublicKey(client *api.Client, key string, opt consulutils.QueryOptions) *rsa.PublicKey {
	val, err := consulutils.GetKV(client, key, opt)
	if err != nil || len(val) == 0 {
		return nil
	}
	block, _ := pem.Decode([]byte(val))
	if block == nil {
		return nil
	}
	if block.Type != "PUBLIC KEY" && block.Type != "RSA PUBLIC KEY" {
		return nil
	}

	pubkeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPublicKey, ok := pubkeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil
	}
	return rsaPublicKey
}

func GetRsaPrivateKey(client *api.Client, key string, opt consulutils.QueryOptions) *rsa.PrivateKey {
	val, err := consulutils.GetKV(client, key, opt)
	if err != nil || len(val) == 0 {
		return nil
	}
	block, _ := pem.Decode([]byte(val))
	if block == nil {
		return nil
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return privateKey
}


func GetSMTPConfig(client *api.Client, path string, opt consulutils.QueryOptions) (*types.SMTPConfig, error) {
	var cfg types.SMTPConfig
	str, err := consulutils.GetKV(client, path, opt)
	if err != nil {
		return nil, ErrCantGetConfig
	}
	err = json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, ErrCantUnmarshalJSON
	}
	return &cfg, nil
}