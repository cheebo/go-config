package utils_test

import (
	"testing"

	"github.com/cheebo/go-config/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type (
	Person struct {
		Name string
		Age  int
	}
	Doc struct {
		ID     int
		Data   *Person
		Events []string
		Tags   *Tags
	}
	Tags  []string
	Human Person
)

func TestStructToStringMap(t *testing.T) {
	a := assert.New(t)

	s := struct {
		Boolean bool
		Integer int
		Float   float64
		String  string
	}{
		Boolean: true,
		Integer: 100,
		Float:   99.9,
		String:  "foo bar",
	}

	m := map[string]interface{}{
		"boolean": s.Boolean,
		"integer": s.Integer,
		"float":   s.Float,
		"string":  s.String,
	}

	result := utils.StructToStringMap(s)

	a.Equal(m, result)
}

func TestStructToStringMap_NestedStruct(t *testing.T) {
	a := assert.New(t)

	s := struct {
		ID   int
		Data struct {
			Key   string
			Value string
		}
	}{
		ID: 1,
		Data: struct {
			Key   string
			Value string
		}{Key: "foo", Value: "bar"},
	}

	m := map[string]interface{}{
		"id": s.ID,
		"data": map[string]interface{}{
			"key":   s.Data.Key,
			"value": s.Data.Value,
		},
	}

	result := utils.StructToStringMap(s)

	a.Equal(m, result)
}

func TestStructToStringMap_NestedPtrStruct(t *testing.T) {
	a := assert.New(t)

	p := Person{
		Name: "John Doe",
		Age:  99,
	}
	d := Doc{
		ID:   10,
		Data: &p,
	}

	m := map[string]interface{}{
		"id": d.ID,
		"data": map[string]interface{}{
			"name": d.Data.Name,
			"age":  d.Data.Age,
		},
		"events": d.Events,
	}

	result := utils.StructToStringMap(d)

	a.Equal(m, result)
}

func TestStructToStringMap_NestedPtrType(t *testing.T) {
	a := assert.New(t)

	d := Doc{
		ID:   10,
		Tags: &Tags{"t1", "t2"},
	}

	m := map[string]interface{}{
		"id":     d.ID,
		"events": d.Events,
		"tags":   *d.Tags,
	}

	result := utils.StructToStringMap(d)

	a.Equal(m, result)
}
