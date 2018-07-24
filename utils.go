package go_config

import (
	"github.com/spf13/cast"
)

func MergeMapWithPath(source map[string]interface{}, sub map[string]interface{}, path []string) error {
	if len(path) == 0 || (len(path) == 1 && path[0] == "") {
		for k, v := range sub {
			source[k] = v
		}
		return nil
	}

	next, ok := source[path[0]]
	if !ok {
		next = map[string]interface{}{}
		source[path[0]] = next
	}

	return MergeMapWithPath(cast.ToStringMap(source[path[0]]), sub, path[1:])
}
