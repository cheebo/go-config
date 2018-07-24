package file_test

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/file"
	"github.com/cheebo/go-config/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

type Config struct {
	AMQP types.AMQPConfig `cfg:"amqp"`
}

func TestJsonFileSource(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	c := go_config.New()
	fileSource, err := file.Source(file.File{Path: "./fixtures/config.json", Type: go_config.JSON, Namespace: ""})
	assert.NoError(err)

	c.UseSource(fileSource)

	assert.Equal("localhost", cfg.AMQP.URL)
	assert.Equal("exch", cfg.AMQP.Exchange)
	assert.Equal("que", cfg.AMQP.Queue)
	assert.Equal("knd", cfg.AMQP.Kind)
	assert.Equal("k", cfg.AMQP.Key)
	assert.Equal(true, cfg.AMQP.Durable)
	assert.Equal(true, cfg.AMQP.AutoDelete)
	assert.Equal(2, int(cfg.AMQP.DeliveryMode))
}
