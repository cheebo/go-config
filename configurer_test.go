package go_config_test

import (
	"github.com/cheebo/go-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SliceConfig struct {
	Names []string `delimiter:","`
	Ips   []string `delimiter:";"`
	V     string   `cfg:"test.v"`
}

func TestConfig_Configure(t *testing.T) {
	assert := assert.New(t)

	var config SliceConfig

	cfg := go_config.New()
	cfg.Use(go_config.EnvironmentSource())
	cfg.Use(go_config.FlagSource())
	err := cfg.Configure(&config)

	assert.NoError(err)
	assert.EqualValues([]string{"name1", "name2", "name3"}, config.Names)
}
