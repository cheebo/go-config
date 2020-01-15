package utils

import (
	"github.com/spf13/cast"
)

// MergeMapWithPath merges sub into source, result in source
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

func Lookup(source map[string]interface{}, key []string) interface{} {
	if len(key) == 0 {
		return source
	}

	next, ok := source[key[0]]
	if ok {
		if len(key) == 1 {
			return next
		}

		// Nested case
		switch next.(type) {
		case map[interface{}]interface{}:
			return Lookup(cast.ToStringMap(next), key[1:])
		case map[string]interface{}:
			return Lookup(next.(map[string]interface{}), key[1:])
		default:
			return nil
		}
	}
	return nil
}
