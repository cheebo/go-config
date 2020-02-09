package go_config_test

import (
	"testing"

	go_config "github.com/cheebo/go-config"
	"github.com/stretchr/testify/assert"
)

func TestGoConfig_GetTypes(t *testing.T) {
	a := assert.New(t)

	const (
		String      = "Foo Bar"
		Int         = 100
		UInt   uint = 200
		Bool        = true
	)
	var (
		Slice = []int{1, 2, 3}
		Map   = map[string]int{"one": 1, "two": 2}
	)

	cfg := go_config.New()
	cfg.SetDefault("string", String)
	cfg.SetDefault("int", Int)
	cfg.SetDefault("uint", UInt)
	cfg.SetDefault("bool", Bool)
	cfg.SetDefault("slice", Slice)
	cfg.SetDefault("map", Map)

	a.Equal(String, cfg.String("string"))
	a.Equal(Int, cfg.Int("int"))
	a.Equal(UInt, cfg.UInt("uint"))
	a.Equal(Bool, cfg.Bool("bool"))
	a.Equal(Slice, cfg.Slice("slice"))
	a.Equal(Map, cfg.StringMap("map"))
}

func TestGoConfig_Sub(t *testing.T) {
	a := assert.New(t)

	const (
		String      = "Foo Bar"
		Int         = 100
		UInt   uint = 200
		Bool        = true
	)
	var (
		Slice = []int{1, 2, 3}
		Map   = map[string]int{"one": 1, "two": 2}
	)

	cfg := go_config.New()
	cfg.SetDefault("foo.string", String)
	cfg.SetDefault("foo.int", Int)
	cfg.SetDefault("foo.uint", UInt)
	cfg.SetDefault("foo.bool", Bool)
	cfg.SetDefault("foo.slice", Slice)
	cfg.SetDefault("foo.map", Map)

	sub := cfg.Sub("foo")

	a.Equal(String, sub.String("string"))
	a.Equal(Int, sub.Int("int"))
	a.Equal(UInt, sub.UInt("uint"))
	a.Equal(Bool, sub.Bool("bool"))
	a.Equal(Slice, sub.Slice("slice"))
	a.Equal(Map, sub.StringMap("map"))
}

func TestGoConfig_SetDefault(t *testing.T) {
	a := assert.New(t)

	const (
		name = "John Doe"
		age  = 99
	)

	type person struct {
		Name string
		Age  int
	}

	p := person{
		Name: name,
		Age:  age,
	}

	cfg := go_config.New()
	cfg.SetDefault("name", name)
	cfg.SetDefault("john.age", age)
	cfg.SetDefault("person", p)

	a.Equal(name, cfg.String("name"))
	a.Equal(age, cfg.Int("john.age"))
	a.Equal(name, cfg.String("person.name"))
	a.Equal(age, cfg.Int("person.age"))

	sub := cfg.Sub("john")
	a.Equal(age, sub.Int("age"))
}

func TestGoConfig_Unmarshal(t *testing.T) {
	a := assert.New(t)

	const (
		name = "John Doe"
		age  = 99
		city = "The City"
		zip  = 10118
	)

	var (
		edu  = []string{"school", "bachelor", "master", "phd"}
		exp  = []int{1, 2, 10, 20}
		jobs = map[string]string{
			"Acme Corp": "Manager",
			"Acme LLC":  "Manager",
		}
		jobExp = map[string]int{
			"acme corp":     3,
			"home business": 1,
		}
		business = map[string][]string{
			"owner":    []string{"home business", "first business"},
			"co-owner": []string{"second corp"},
		}
	)

	type person struct {
		Name    string
		Age     int
		Address struct {
			City string
			Zip  int
		}
		Jobs       map[string]string
		JobExp     map[string]int
		Business   map[string][]string
		Business2  map[string]interface{}
		Education  []string
		Experience []int
		Travel     []interface{}
	}

	cfg := go_config.New()
	cfg.SetDefault("person.name", name)
	cfg.SetDefault("person.age", age)
	cfg.SetDefault("person.address.city", city)
	cfg.SetDefault("person.address.zip", zip)
	cfg.SetDefault("person.education", edu)
	cfg.SetDefault("person.experience", exp)
	cfg.SetDefault("person.travel", edu)
	cfg.SetDefault("person.jobs", jobs)
	cfg.SetDefault("person.jobexp", jobExp)
	cfg.SetDefault("person.business", business)
	cfg.SetDefault("person.business2", business)

	p := person{}

	err := cfg.Unmarshal(&p, "person")
	a.NoError(err)
	a.Equal(name, p.Name)
	a.Equal(age, p.Age)
	a.Equal(city, p.Address.City)
	a.Equal(zip, p.Address.Zip)
	a.Equal(edu, p.Education)
	a.Equal(exp, p.Experience)
	a.Equal(jobs, p.Jobs)
	a.Equal(business, p.Business)
	a.Equal(map[string]interface{}{}, p.Business2)
	a.Equal(jobExp, p.JobExp)
}

func TestGoConfig_GetUnset(t *testing.T) {
	a := assert.New(t)

	cfg := go_config.New()
	v := cfg.Get("unset")
	a.Nil(v)
}
