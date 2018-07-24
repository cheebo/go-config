package go_config_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/cheebo/go-config"
	"strings"
	"testing"
)

var (
	base = map[string]interface{}{
		"name": "Jhon Doe",
		"age":  30,
	}

	m = map[string]interface{}{
		"addr": "localhost",
		"port": 3000,
	}

	dst = map[string]interface{}{
		"name": "Jhon Doe",
		"age":  30,
		"db": map[string]interface{}{
			"one": map[string]interface{}{
				"addr": "localhost",
				"port": 3000,
			},
		},
	}
)

func TestMergeMap(t *testing.T) {
	assert := assert.New(t)

	err := go_config.MergeMapWithPath(base, m, strings.Split("db.one", "."))

	assert.NoError(err)
	assert.Equal(dst, base)
}
