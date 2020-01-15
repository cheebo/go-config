package file_test

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/internal/reader"
	"github.com/cheebo/go-config/sources/file"
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
	AMQP AMQPConfig
}

func TestJsonFileSource(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	c := go_config.New()
	fileSource, err := file.Source(file.File{Path: "./fixtures/config.json", Type: reader.JSON, Namespace: ""})
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
