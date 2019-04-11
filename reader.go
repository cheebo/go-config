package go_config

import (
	"bytes"
	"encoding/json"
	"fmt"
	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v2"
	"io"
)

type ConfigType int

type ConfigParseError struct {
	err error
}

const (
	JSON ConfigType = iota
	YAML
	TOML
)

func (pe ConfigParseError) Error() string {
	return fmt.Sprintf("Config parsing error: %s", pe.err.Error())
}

func ReadConfig(in io.Reader, ct ConfigType, c map[string]interface{}) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(in)
	if err != nil {
		return err
	}

	switch ct {
	case JSON:
		decoder := json.NewDecoder(buf)
		decoder.UseNumber()
		if err := decoder.Decode(&c); err != nil {
			return ConfigParseError{err}
		}
	case YAML:
		if err := yaml.Unmarshal(buf.Bytes(), &c); err != nil {
			return ConfigParseError{err}
		}
	case TOML:
		tree, err := toml.LoadReader(buf)
		if err != nil {
			return ConfigParseError{err}
		}
		tomlMap := tree.ToMap()
		for k, v := range tomlMap {
			c[k] = v
		}
	}

	return nil
}
