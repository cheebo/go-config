package reader_test

import (
	"bytes"
	"fmt"
	"github.com/cheebo/go-config/internal/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = struct {
	Human bool
	Name string
	Age int
	Height float64
	Roles []interface{}
	Map map[string]interface{}
}{
	Human: true,
	Name: "John",
	Age: 50,
	Height: 6.0,
	Roles: []interface{}{"admin", "manager"},
	Map: map[string]interface{}{
		"planet": "Earth",
		"system": "sol1",
	},
}

func TestReadConfig_YAML(t *testing.T) {
	assert := assert.New(t)

	var yamlExample = []byte(`
human: true
name: John
age: 50
height: 6.0
roles:
- admin
- manager
map:
  planet: Earth
  system: sol1
`)
	cfg := map[string]interface{}{}
	err := reader.ReadConfig(bytes.NewBuffer(yamlExample), reader.YAML, cfg)

	assert.NoError(err)

	assert.Equal(data.Human, cfg["human"])
	assert.Equal(data.Name, cfg["name"])
	assert.Equal(data.Age, cfg["age"])
	assert.Equal(data.Height, cfg["height"])
	assert.Equal(data.Roles, cfg["roles"])
}

func TestReadConfig_JSON(t *testing.T) {
	assert := assert.New(t)

	var jsonExample = []byte(`{
"human": true,
"name": "John",
"age": 50,
"height": 6.0,
"roles": ["admin", "manager"],
"map": {"planet": "Earth", "system": "sol1"}
}`)
	cfg := map[string]interface{}{}
	err := reader.ReadConfig(bytes.NewBuffer(jsonExample), reader.JSON, cfg)

	assert.NoError(err)

	assert.Equal(data.Human, cfg["human"])
	assert.Equal(data.Name, cfg["name"])
	assert.Equal(fmt.Sprintf("%d", data.Age), fmt.Sprintf("%v", cfg["age"]))
	assert.Equal(fmt.Sprintf("%.1f", data.Height), fmt.Sprintf("%v", cfg["height"]))
	assert.Equal(data.Roles, cfg["roles"])
	assert.Equal(data.Map, cfg["map"])
}

func TestReadConfig_TOML(t *testing.T) {
	assert := assert.New(t)

	var jsonExample = []byte(`
human = true
name = "John"
age = 50.0
height = 6.0
roles = [
  "admin",
  "manager"
]

[map]
planet = "Earth"
system = "sol1"
`)
	cfg := map[string]interface{}{}
	err := reader.ReadConfig(bytes.NewBuffer(jsonExample), reader.TOML, cfg)

	assert.NoError(err)

	assert.Equal(data.Human, cfg["human"])
	assert.Equal(data.Name, cfg["name"])
	assert.Equal(float64(data.Age), cfg["age"])
	assert.Equal(data.Height, cfg["height"])
	assert.Equal(data.Roles, cfg["roles"])
	assert.Equal(data.Map, cfg["map"])
}
