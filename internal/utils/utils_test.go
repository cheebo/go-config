package utils_test

import (
	"github.com/cheebo/go-config/internal/utils"
	"github.com/stretchr/testify/assert"

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

	dst2 = map[string]interface{}{
		"name": "Jhon Doe",
		"age":  30,
		"addr": "localhost",
		"port": 3000,
	}
)

func TestMergeMap_Path(t *testing.T) {
	assert := assert.New(t)

	b := map[string]interface{}{}
	for k, v := range base {
		b[k] = v
	}

	err := utils.MergeMapWithPath(b, m, strings.Split("db.one", "."))

	assert.NoError(err)
	assert.Equal(dst, b)
}

func TestMergeMap_EmptyPath(t *testing.T) {
	assert := assert.New(t)

	b := map[string]interface{}{}
	for k, v := range base {
		b[k] = v
	}
	err := utils.MergeMapWithPath(b, m, []string{})

	assert.NoError(err)
	assert.Equal(dst2, b)
}