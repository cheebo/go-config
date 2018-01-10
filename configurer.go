package go_config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrNotAStructPtr = errors.New("expects pointer to a struct")
)

type Config interface {
	Use(sources ...Source)
	Configure(v interface{}) error
	Usage() map[string]string
}

// Source: implement this interface to get configurations from sources like env, flag, file, kv-store etc
type Source interface {
	Init(variables map[string]*Variable) error
	Int(name string) (int, error)
	Float(name string) (float64, error)
	UInt(name string) (uint, error)
	String(name string) (string, error)
	Bool(name string) (bool, error)
}

type ConfigSource interface {
	Load() error
}

type config struct {
	sources   []Source
	variables map[string]*Variable
}

func New() Config {
	return &config{
		sources:   []Source{},
		variables: make(map[string]*Variable),
	}
}

func (self *config) Use(sources ...Source) {
	self.sources = append(self.sources, sources...)
}

func (self *config) Usage() map[string]string {
	usage := make(map[string]string)
	for k, v := range self.variables {
		usage[k] = v.Description
	}
	return usage
}

func (self *config) Configure(v interface{}) error {
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return ErrNotAStructPtr
	}
	ref := ptr.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotAStructPtr
	}

	self.setup(v, "")

	for _, src := range self.sources {
		err := src.Init(self.variables)
		if err != nil {
			return err
		}
	}

	return self.fillData()
}

func (self *config) setup(v interface{}, parent string) error {
	refVal := reflect.ValueOf(v)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	if refVal.Kind() != reflect.Struct {
		return nil
	}

	refType := reflect.TypeOf(refVal.Interface())

	for i := 0; i < refVal.NumField(); i++ {
		field := refType.Field(i)
		refField := refVal.Field(i)

		name := field.Name
		tagName, _ := parseTag(field.Tag.Get("cfg"))
		if len(tagName) > 0 {
			name = tagName
		}
		if len(parent) > 0 {
			name = parent + "." + name
		}

		if refField.Kind() == reflect.Ptr {
			if refField.IsNil() {
				refField = reflect.New(refField.Type().Elem())
				refVal.Field(i).Set(refField)
				refField = refField.Elem()
			} else {
				refField = refField.Elem()
			}
		}

		if refField.Kind() == reflect.Struct {
			self.setup(refField.Addr().Interface(), name)
			self.variables[name] = &Variable{
				Name:        name,
				Description: field.Tag.Get("description"),
				Tag:         field.Tag,
				Type:        refField.Type(),
			}
			continue
		}

		if !refField.CanSet() {
			continue
		}

		self.variables[name] = &Variable{
			Name:        name,
			Description: field.Tag.Get("description"),
			Def:         reflect.Zero(refField.Type()),
			Set:         refField.Set,
			Tag:         field.Tag,
			Type:        refField.Type(),
		}

	}
	return nil
}

func (self *config) fillData() error {
	for _, val := range self.variables {
		changed := false

		for _, src := range self.sources {

			switch val.Type.Kind() {
			case reflect.Int:
				s, err := src.Int(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}

				val.set(s)

			case reflect.Uint:
				s, err := src.UInt(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}

				val.set(s)

			case reflect.Float64:
				s, err := src.Float(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}

				val.set(s)

			case reflect.String:
				s, err := src.String(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}

				val.set(s)

			case reflect.Bool:
				s, err := src.Bool(val.Name)
				if err != nil {
					continue
				}
				if reflect.Zero(val.Type).Interface() == reflect.ValueOf(&s).Elem().Interface() {
					continue
				}

				val.set(s)
			}

			changed = true
		}
		if !changed {
			val.set(val.Def.Interface())
		}
	}
	return nil
}

func parseTag(tag string) (string, []string) {
	opts := strings.Split(tag, ",")
	return opts[0], opts[1:]
}

// Variable Routines
type Variable struct {
	Name        string
	Def         reflect.Value
	Description string
	Set         func(x reflect.Value)
	Tag         reflect.StructTag
	Type        reflect.Type
}

func (v Variable) String() string {
	return fmt.Sprintf("%v[%v] %v", v.Name, v.Type.Kind(), v.Description)
}

func (v *Variable) set(value interface{}) {
	if v.Type.Kind() == reflect.Struct {
		return
	}
	v.Set(reflect.ValueOf(value).Convert(v.Type))
}
