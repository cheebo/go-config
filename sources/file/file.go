package file

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/internal/reader"
	"github.com/cheebo/go-config/internal/utils"
	"github.com/spf13/cast"
)

type File struct {
	Path      string
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

		var cfgType reader.ConfigType
		switch ext := strings.ToLower(path.Ext(f.Path)); ext {
		case "json":
			cfgType = reader.JSON
		case "toml":
			cfgType = reader.TOML
		case "yaml":
			fallthrough
		case "yml":
			cfgType = reader.YAML
		default:
			return nil, errors.New("unsupported file extensions: " + ext)
		}

		m := map[string]interface{}{}
		err = reader.ReadConfig(bytes.NewBuffer(data), cfgType, m)
		if err != nil {
			return nil, err
		}

		utils.MergeMapWithPath(config, m, strings.Split(f.Namespace, "."))
	}

	return &file{
		fs:   fs,
		data: config,
	}, nil
}

func (f *file) Get(key string) interface{} {
	return utils.Lookup(f.data, f.key(key))
}

func (f *file) Bool(key string) (bool, error) {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return false, go_config.NoVariablesInitialised
	}
	return cast.ToBoolE(val)
}

func (f *file) Float(key string) (float64, error) {
	val := utils.Lookup(f.data, f.key(key))
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
	val := utils.Lookup(f.data, f.key(key))
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

func (f *file) Int8(key string) (int8, error) {
	val := utils.Lookup(f.data, f.key(key))
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

	return cast.ToInt8E(v)
}

func (f *file) Int32(key string) (int32, error) {
	val := utils.Lookup(f.data, f.key(key))
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

	return cast.ToInt32E(v)
}

func (f *file) Int64(key string) (int64, error) {
	val := utils.Lookup(f.data, f.key(key))
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

	return cast.ToInt64E(v)
}

func (f *file) UInt(key string) (uint, error) {
	val := utils.Lookup(f.data, f.key(key))
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

func (f *file) UInt32(key string) (uint32, error) {
	val := utils.Lookup(f.data, f.key(key))
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

	return cast.ToUint32E(v)
}

func (f *file) UInt64(key string) (uint64, error) {
	val := utils.Lookup(f.data, f.key(key))
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

	return cast.ToUint64E(v)
}

func (f *file) Slice(key, delimiter string) ([]interface{}, error) {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return []interface{}{}, go_config.NoVariablesInitialised
	}
	return cast.ToSliceE(val)
}

func (f *file) SliceInt(key string) ([]int, error) {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return []int{}, go_config.NoVariablesInitialised
	}
	return cast.ToIntSliceE(val)
}

func (f *file) SliceString(key string) ([]string, error) {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return []string{}, go_config.NoVariablesInitialised
	}
	return cast.ToStringSliceE(val)
}

func (f *file) String(key string) (string, error) {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return "", go_config.NoVariablesInitialised
	}
	return cast.ToStringE(val)
}

func (f *file) StringMap(key string) map[string]interface{} {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return map[string]interface{}{}
	}
	return cast.ToStringMap(val)
}

func (f *file) StringMapInt(key string) map[string]int {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return map[string]int{}
	}
	return cast.ToStringMapInt(val)
}

func (f *file) StringMapSliceString(key string) map[string][]string {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return map[string][]string{}
	}
	return cast.ToStringMapStringSlice(val)
}

func (f *file) StringMapString(key string) map[string]string {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return map[string]string{}
	}
	return cast.ToStringMapString(val)
}

func (f *file) IsSet(key string) bool {
	val := utils.Lookup(f.data, f.key(key))
	if val == nil {
		return false
	}
	return true
}

func (f *file) key(key string) []string {
	return strings.Split(key, ".")
}
