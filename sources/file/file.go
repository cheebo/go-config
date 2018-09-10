package file

import (
	"bytes"
	"encoding/json"
	"github.com/cheebo/go-config"
	"github.com/spf13/cast"
	"io/ioutil"
	"strings"
)

type File struct {
	Path      string
	Type      go_config.ConfigType
	Namespace string
}

type file struct {
	fs   []File
	ns   string
	data map[string]interface{}
}

func Source(fs ...File) (go_config.Source, error) {
	config := map[string]interface{}{}

	for _, f := range fs {
		data, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, err
		}

		m := map[string]interface{}{}
		err = go_config.ReadConfig(bytes.NewBuffer(data), f.Type, m)
		if err != nil {
			return nil, err
		}

		go_config.MergeMapWithPath(config, m, strings.Split(f.Namespace, "."))
	}

	return &file{
		fs:   fs,
		data: config,
	}, nil
}

func (f *file) Get(key string) interface{} {
	return go_config.Lookup(f.data, f.key(key))
}

func (f *file) Bool(key string) (bool, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (f *file) Float(key string) (float64, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}

	var v interface{}
	switch val.(type) {
	case json.Number:
		v = val.(json.Number).String()
		break
	default:
		v = val
	}

	return cast.ToFloat64E(v)
}

func (f *file) Int(key string) (int, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}

	var v interface{}
	switch val.(type) {
	case json.Number:
		v = val.(json.Number).String()
		break
	default:
		v = val
	}

	return cast.ToIntE(v)
}

func (f *file) UInt(key string) (uint, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return 0, go_config.NoVariablesInitialised
	}

	var v interface{}
	switch val.(type) {
	case json.Number:
		v = val.(json.Number).String()
		break
	default:
		v = val
	}

	return cast.ToUintE(v)
}

func (f *file) Slice(key, delimiter string) ([]interface{}, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(val)
}

func (f *file) String(key string) (string, error) {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (f *file) StringMap(key string) map[string]interface{} {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (f *file) IsSet(key string) bool {
	val := go_config.Lookup(f.data, f.key(key))
	if val == nil {
		return false
	}
	return true
}

func (f *file) key(key string) []string {
	return strings.Split(key, ".")
}
