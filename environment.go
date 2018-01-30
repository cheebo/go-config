package go_config

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

type env struct{}

func EnvironmentSource() Source {
	return env{}
}

func (self env) Init(vals map[string]*Variable) error {
	return nil
}

func (self env) Int(name string) (int, error) {
	val := os.Getenv(self.name(name))
	return strconv.Atoi(val)
}

func (self env) Float(name string) (float64, error) {
	val, err := strconv.ParseFloat(os.Getenv(self.name(name)), 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (self env) UInt(name string) (uint, error) {
	val, err := strconv.ParseUint(os.Getenv(self.name(name)), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func (self env) String(name string) (string, error) {
	return os.Getenv(self.name(name)), nil
}

func (self env) Bool(name string) (bool, error) {
	val := os.Getenv(self.name(name))
	return strconv.ParseBool(val)
}

func (self env) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	src := os.Getenv(self.name(name))
	return parseSlice(src, delimiter, kind)
}

func (self env) name(name string) string {
	return strings.Replace(strings.ToUpper(name), ".", "_", -1)
}
