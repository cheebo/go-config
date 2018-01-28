package go_config_test

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

type ConfigConsul struct {
	AMQP types.AMQPConfig `cfg:"amqp" consul:"/test/amqp"`
}

func TestConsulSource(t *testing.T) {
	assert := assert.New(t)
	cfg := ConfigConsul{}
	c := go_config.New()
	c.Use(go_config.ConsulSource("/s2w", types.ConsulConfig{
		Addr: "localhost:8500", Scheme: "http",
	}))
	err := c.Configure(&cfg)

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
