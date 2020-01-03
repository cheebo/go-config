package consul_test

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/consul"
	"github.com/stretchr/testify/assert"

	"testing"
)

type AMQPConfig struct {
	URL          string `json:"url"`
	Exchange     string `json:"exchange"`
	Queue        string `json:"queue"`
	Kind         string `json:"kind"`
	Key          string `json:"key"`
	Durable      bool   `json:"durable"`
	AutoDelete   bool   `json:"auto_delete"`
	DeliveryMode uint   `json:"delivery_mode"`
}

type Config struct {
	AMQP AMQPConfig `cfg:"amqp" consul:"/test/amqp"`
}

type ConfigRsaKey struct {
	Key string `consul:"/rsa/public"`
}

func TestConsulSource(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	c := go_config.New()

	src, err := consul.Source(consul.ConsulConfig{
		Addr: "localhost:8500", Scheme: "http",
	})
	assert.NoError(err)

	c.UseSource(src)

	err = c.Unmarshal(cfg, "amqp")
	assert.NoError(err)

	assert.Equal("localhost", cfg.AMQP.URL)
	assert.Equal("exch", cfg.AMQP.Exchange)
	assert.Equal("que", cfg.AMQP.Queue)
	assert.Equal("knd", cfg.AMQP.Kind)
	assert.Equal("k", cfg.AMQP.Key)
	assert.Equal(true, cfg.AMQP.Durable)
	assert.Equal(true, cfg.AMQP.AutoDelete)
	assert.Equal(2, int(cfg.AMQP.DeliveryMode))
}

func TestConsulSource2(t *testing.T) {
	assert := assert.New(t)
	cfg := ConfigRsaKey{}
	c := go_config.New()

	src, err := consul.Source(consul.ConsulConfig{
		Addr: "localhost:8500", Scheme: "http",
	})
	assert.NoError(err)

	c.UseSource(src)
	err = c.Unmarshal(&cfg, "rsa.public")

	assert.NoError(err)

	assert.Equal("RSA PUBLIC KEY", cfg.Key)
}
