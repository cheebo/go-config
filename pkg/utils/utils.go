package utils

import (
	"reflect"
	"strings"
)

func StructToStringMap(v interface{}) map[string]interface{} {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	m := map[string]interface{}{}
	//
	typeOfVal := reflect.TypeOf(val.Interface())
	for i := 0; i < val.NumField(); i++ {
		var fieldValue interface{}

		field := val.Field(i)

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			fieldValue = StructToStringMap(field.Interface())
		} else {
			fieldValue = field.Interface()
		}

		m[strings.ToLower(typeOfVal.Field(i).Name)] = fieldValue
	}

	return m
}
