package go_config_test

import (
	"testing"

	go_config "github.com/cheebo/go-config"
	"github.com/cheebo/go-config/pkg/errors"
	"github.com/cheebo/go-config/pkg/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	keyBool                 = "key_bool"
	keyFloat                = "key_float"
	keyInt                  = "key_int"
	keyInt8                 = "key_int8"
	keyInt32                = "key_int32"
	keyInt64                = "key_int64"
	keySlice                = "key_slice"
	keySliceInt             = "key_sliceInt"
	keySliceString          = "key_sliceString"
	keyString               = "key_string"
	keyStringMap            = "key_stringMap"
	keyStringMapInt         = "key_stringMapInt"
	keyStringMapSliceString = "key_stringMapSliceString"
	keyStringMapString      = "key_stringMapString"
	keyUInt                 = "key_uint"
	keyUInt32               = "key_uint32"
	keyUInt64               = "key_uint64"
)

const (
	valBool           = true
	valFloat  float64 = 64
	valInt    int     = 128
	valInt8   int8    = 8
	valInt32  int32   = 32
	valInt64  int64   = 64
	valUInt   uint    = 256
	valUInt32 uint32  = 512
	valUInt64 uint64  = 1024

	defBool           = true
	defFloat  float64 = 640
	defInt    int     = 1280
	defInt8   int8    = 80
	defInt32  int32   = 320
	defInt64  int64   = 640
	defUInt   uint    = 2560
	defUInt32 uint32  = 5120
	defUInt64 uint64  = 10240
)

var (
	valSlice                = []interface{}{1, 2, 3}
	valSliceInt             = []int{1, 3, 2}
	valSliceString          = []string{"_foo_", "_bar_", "_baz_"}
	valString               = "hello world"
	valStringMap            = map[string]interface{}{"int8": 256}
	valStringMapInt         = map[string]int{"int": 256}
	valStringMapSliceString = map[string][]string{"int": []string{"_foo", "_bar", "_baz"}}
	valStringMapString      = map[string]string{"key": "value"}
)

func TestConfig_GetFromSource(t *testing.T) {
	t.Run("config.Bool", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Bool(keyBool).Return(valBool, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valBool, cfg.Bool(keyBool))
	})

	t.Run("config.Float", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Float(keyFloat).Return(valFloat, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valFloat, cfg.Float(keyFloat))
	})

	t.Run("config.Int", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int(keyInt).Return(valInt, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valInt, cfg.Int(keyInt))
	})

	t.Run("config.Int8", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int8(keyInt8).Return(valInt8, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valInt8, cfg.Int8(keyInt8))
	})

	t.Run("config.Int32", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int32(keyInt32).Return(valInt32, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valInt32, cfg.Int32(keyInt32))
	})

	t.Run("config.Int64", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int64(keyInt64).Return(valInt64, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valInt64, cfg.Int64(keyInt64))
	})

	t.Run("config.Slice", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Slice(keySlice).Return(valSlice, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valSlice, cfg.Slice(keySlice))
	})

	t.Run("config.SliceInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().SliceInt(keySliceInt).Return(valSliceInt, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valSliceInt, cfg.SliceInt(keySliceInt))
	})

	t.Run("config.SliceString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Slice(keySlice).Return(valSlice, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valSlice, cfg.Slice(keySlice))
	})

	t.Run("config.String", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().String(keyString).Return(valString, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valString, cfg.String(keyString))
	})

	t.Run("config.StringMap", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMap(keyStringMap).Return(valStringMap).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valStringMap, cfg.StringMap(keyStringMap))
	})

	t.Run("config.StringMapInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapInt(keyStringMapInt).Return(valStringMapInt).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valStringMapInt, cfg.StringMapInt(keyStringMapInt))
	})

	t.Run("config.StringMapSliceString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapSliceString(keyStringMapSliceString).Return(valStringMapSliceString).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valStringMapSliceString, cfg.StringMapSliceString(keyStringMapSliceString))
	})

	t.Run("config.StringMapString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapString(keyStringMapString).Return(valStringMapString).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valStringMapString, cfg.StringMapString(keyStringMapString))
	})

	t.Run("config.UInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt(keyUInt).Return(valUInt, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valUInt, cfg.UInt(keyUInt))
	})

	t.Run("config.UInt32", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt32(keyUInt32).Return(valUInt32, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valUInt32, cfg.UInt32(keyUInt32))
	})

	t.Run("config.UInt64", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt64(keyUInt64).Return(valUInt64, nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		assert.Equal(t, valUInt64, cfg.UInt64(keyUInt64))
	})
}

func TestConfig_SetDefault(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("config.Bool", func(t *testing.T) {
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Bool(keyBool).Return(false, errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyBool, valBool)
		assert.Equal(t, valBool, cfg.Bool(keyBool))
	})

	t.Run("config.Float", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Float(keyFloat).Return(float64(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyFloat, valFloat)
		assert.Equal(t, valFloat, cfg.Float(keyFloat))
	})

	t.Run("config.Int", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int(keyInt).Return(0, errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyInt, valInt)
		assert.Equal(t, valInt, cfg.Int(keyInt))
	})

	t.Run("config.Int8", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int8(keyInt8).Return(int8(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyInt8, valInt8)
		assert.Equal(t, valInt8, cfg.Int8(keyInt8))
	})

	t.Run("config.Int32", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int32(keyInt32).Return(int32(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyInt32, valInt32)
		assert.Equal(t, valInt32, cfg.Int32(keyInt32))
	})

	t.Run("config.Int64", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Int64(keyInt64).Return(int64(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyInt64, valInt64)
		assert.Equal(t, valInt64, cfg.Int64(keyInt64))
	})

	t.Run("config.Slice", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().Slice(keySlice).Return(nil, errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keySlice, valSlice)
		assert.Equal(t, valSlice, cfg.Slice(keySlice))
	})

	t.Run("config.SliceInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().SliceInt(keySliceInt).Return(nil, errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keySliceInt, valSliceInt)
		assert.Equal(t, valSliceInt, cfg.SliceInt(keySliceInt))
	})

	t.Run("config.SliceString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().SliceString(keySliceString).Return(nil, errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keySliceString, valSliceString)
		assert.Equal(t, valSliceString, cfg.SliceString(keySliceString))
	})

	t.Run("config.String", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().String(keyString).Return("", errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyString, valString)
		assert.Equal(t, valString, cfg.String(keyString))
	})

	t.Run("config.StringMap", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMap(keyStringMap).Return(nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyStringMap, valStringMap)
		assert.Equal(t, valStringMap, cfg.StringMap(keyStringMap))
	})

	t.Run("config.StringMapInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapInt(keyStringMapInt).Return(nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyStringMapInt, valStringMapInt)
		assert.Equal(t, valStringMapInt, cfg.StringMapInt(keyStringMapInt))
	})

	t.Run("config.StringMapSliceString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapSliceString(keyStringMapSliceString).Return(nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyStringMapSliceString, valStringMapSliceString)
		assert.Equal(t, valStringMapSliceString, cfg.StringMapSliceString(keyStringMapSliceString))
	})

	t.Run("config.StringMapString", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().StringMapString(keyStringMapString).Return(nil).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyStringMapString, valStringMapString)
		assert.Equal(t, valStringMapString, cfg.StringMapString(keyStringMapString))
	})

	t.Run("config.UInt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt(keyUInt).Return(uint(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyUInt, valUInt)
		assert.Equal(t, valUInt, cfg.UInt(keyUInt))
	})

	t.Run("config.UInt32", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt32(keyUInt32).Return(uint32(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyUInt32, valUInt32)
		assert.Equal(t, valUInt32, cfg.UInt32(keyUInt32))
	})

	t.Run("config.UInt64", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		source := mocks.NewMockSource(ctrl)
		source.EXPECT().UInt64(keyUInt64).Return(uint64(0), errors.NoVariablesInitialised).AnyTimes()
		cfg := go_config.New()
		cfg.UseSource(source)
		cfg.SetDefault(keyUInt64, valUInt64)
		assert.Equal(t, valUInt64, cfg.UInt64(keyUInt64))
	})
}

func TestGoConfig_DefinedInSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	source := mocks.NewMockSource(ctrl)
	namespace := "namespace."
	// fields
	source.EXPECT().Bool(namespace+keyBool).Return(valBool, nil).AnyTimes()
	source.EXPECT().Float(namespace+keyFloat).Return(valFloat, nil).AnyTimes()
	source.EXPECT().Int(namespace+keyInt).Return(valInt, nil).AnyTimes()
	source.EXPECT().Int8(namespace+keyInt8).Return(valInt8, nil).AnyTimes()
	source.EXPECT().Int32(namespace+keyInt32).Return(valInt32, nil).AnyTimes()
	source.EXPECT().Int64(namespace+keyInt64).Return(valInt64, nil).AnyTimes()
	source.EXPECT().Slice(namespace+keySlice).Return(valSlice, nil).AnyTimes()
	source.EXPECT().SliceInt(namespace+keySliceInt).Return(valSliceInt, nil).AnyTimes()
	source.EXPECT().SliceString(namespace+keySliceString).Return(valSliceString, nil).AnyTimes()
	source.EXPECT().String(namespace+keyString).Return(valString, nil).AnyTimes()
	source.EXPECT().StringMap(namespace + keyStringMap).Return(valStringMap).AnyTimes()
	source.EXPECT().StringMapInt(namespace + keyStringMapInt).Return(valStringMapInt).AnyTimes()
	source.EXPECT().StringMapSliceString(namespace + keyStringMapSliceString).Return(valStringMapSliceString).AnyTimes()
	source.EXPECT().StringMapString(namespace + keyStringMapString).Return(valStringMapString).AnyTimes()
	source.EXPECT().UInt(namespace+keyUInt).Return(valUInt, nil).AnyTimes()
	source.EXPECT().UInt32(namespace+keyUInt32).Return(valUInt32, nil).AnyTimes()
	source.EXPECT().UInt64(namespace+keyUInt64).Return(valUInt64, nil).AnyTimes()
	source.EXPECT().StringMap(namespace + keyStringMap).Return(valStringMap).AnyTimes()
	source.EXPECT().StringMapInt(namespace + keyStringMapInt).Return(valStringMapInt).AnyTimes()
	source.EXPECT().StringMapSliceString(namespace + keyStringMapSliceString).Return(valStringMapSliceString).AnyTimes()
	source.EXPECT().StringMapString(namespace + keyStringMapString).Return(valStringMapString).AnyTimes()

	cfg := go_config.New()
	cfg.UseSource(source)
	sub := cfg.Sub("namespace")

	assert.Equal(t, valBool, sub.Bool(keyBool))
	assert.Equal(t, valFloat, sub.Float(keyFloat))
	assert.Equal(t, valInt, sub.Int(keyInt))
	assert.Equal(t, valInt8, sub.Int8(keyInt8))
	assert.Equal(t, valInt32, sub.Int32(keyInt32))
	assert.Equal(t, valInt64, sub.Int64(keyInt64))
	assert.Equal(t, valSlice, sub.Slice(keySlice))
	assert.Equal(t, valSliceInt, sub.SliceInt(keySliceInt))
	assert.Equal(t, valSliceString, sub.SliceString(keySliceString))
	assert.Equal(t, valString, sub.String(keyString))
	assert.Equal(t, valStringMap, sub.StringMap(keyStringMap))
	assert.Equal(t, valStringMapInt, sub.StringMapInt(keyStringMapInt))
	assert.Equal(t, valStringMapSliceString, sub.StringMapSliceString(keyStringMapSliceString))
	assert.Equal(t, valStringMapString, sub.StringMapString(keyStringMapString))
	assert.Equal(t, valUInt, sub.UInt(keyUInt))
	assert.Equal(t, valUInt32, sub.UInt32(keyUInt32))
	assert.Equal(t, valUInt64, sub.UInt64(keyUInt64))

}

//func TestGoConfig_Unmarshal(t *testing.T) {
//	a := assert.New(t)
//
//	const (
//		name = "John Doe"
//		age  = 99
//		city = "The City"
//		zip  = 10118
//	)
//
//	var (
//		edu  = []string{"school", "bachelor", "master", "phd"}
//		exp  = []int{1, 2, 10, 20}
//		jobs = map[string]string{
//			"Acme Corp": "Manager",
//			"Acme LLC":  "Manager",
//		}
//		jobExp = map[string]int{
//			"acme corp":     3,
//			"home business": 1,
//		}
//		business = map[string][]string{
//			"owner":    []string{"home business", "first business"},
//			"co-owner": []string{"second corp"},
//		}
//	)
//
//	type person struct {
//		Name    string
//		Age     int
//		Address struct {
//			City string
//			Zip  int
//		}
//		Jobs       map[string]string
//		JobExp     map[string]int
//		Business   map[string][]string
//		Business2  map[string]interface{}
//		Education  []string
//		Experience []int
//		Travel     []interface{}
//	}
//
//	cfg := go_config.New()
//	cfg.SetDefault("person.name", name)
//	cfg.SetDefault("person.age", age)
//	cfg.SetDefault("person.address.city", city)
//	cfg.SetDefault("person.address.zip", zip)
//	cfg.SetDefault("person.education", edu)
//	cfg.SetDefault("person.experience", exp)
//	cfg.SetDefault("person.travel", edu)
//	cfg.SetDefault("person.jobs", jobs)
//	cfg.SetDefault("person.jobexp", jobExp)
//	cfg.SetDefault("person.business", business)
//	cfg.SetDefault("person.business2", business)
//
//	p := person{}
//
//	err := cfg.Unmarshal(&p, "person")
//	a.NoError(err)
//	a.Equal(name, p.Name)
//	a.Equal(age, p.Age)
//	a.Equal(city, p.Address.City)
//	a.Equal(zip, p.Address.Zip)
//	a.Equal(edu, p.Education)
//	a.Equal(exp, p.Experience)
//	a.Equal(jobs, p.Jobs)
//	a.Equal(business, p.Business)
//	a.Equal(map[string]interface{}{}, p.Business2)
//	a.Equal(jobExp, p.JobExp)
//}

//func TestGoConfig_GetUnset(t *testing.T) {
//	a := assert.New(t)
//
//	cfg := go_config.New()
//	v := cfg.Get("unset")
//	a.Nil(v)
//}
