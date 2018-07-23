package go_config_test

import (
	"bytes"
	goconfig "github.com/cheebo/go-config"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig_YAML(t *testing.T) {
	assert := assert.New(t)

	var yamlExample = []byte(`
admin: true
email: admin@example.com
idle: 50
roles:
- admin
- manager
- client
db:
  addr: localhost
  port: 3306
`)
	cfg := map[string]interface{}{}
	err := goconfig.ReadConfig(bytes.NewBuffer(yamlExample), goconfig.YAML, cfg)

	assert.NoError(err)

	spew.Dump(cfg)
}

func TestReadConfig_JSON(t *testing.T) {
	assert := assert.New(t)

	var jsonExample = []byte(`{
	"admin": true,
	"email": "admin@example.com",
	"idle": 50,
	"roles": ["admin","manager","client"],
	"db": {
	  "addr": "localhost",
	  "port": 3306
	}
}`)
	cfg := map[string]interface{}{}
	err := goconfig.ReadConfig(bytes.NewBuffer(jsonExample), goconfig.JSON, cfg)

	assert.NoError(err)

	spew.Dump(cfg)
}

func TestReadConfig_TOML(t *testing.T) {
	assert := assert.New(t)

	var jsonExample = []byte(`
admin = true
roles = ["admin", "manager", "client"]
idle = 50

[db]
addr = "localhost"
port = 3306
`)
	cfg := map[string]interface{}{}
	err := goconfig.ReadConfig(bytes.NewBuffer(jsonExample), goconfig.TOML, cfg)

	assert.NoError(err)

	spew.Dump(cfg)
}
