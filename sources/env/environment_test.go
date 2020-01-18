package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cheebo/go-config/sources/env"
	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	a := assert.New(t)
	var (
		True                = true
		False               = false
		Int         int     = 1
		Int8        int8    = 2
		Int32       int32   = 3
		Int64       int64   = 4
		UInt        uint    = 6
		UInt32      uint32  = 7
		UInt64      uint64  = 8
		Float       float64 = 3.14
		String              = "Foo bar"
		Slice               = "1,a"
		SliceInt            = "1,2,3"
		SliceString         = "a,b,c,1,2,3"
	)

	os.Setenv("TRUE", "true")
	os.Setenv("FALSE", "false")
	os.Setenv("INT", fmt.Sprintf("%d", Int))
	os.Setenv("INT8", fmt.Sprintf("%d", Int8))
	os.Setenv("INT32", fmt.Sprintf("%d", Int32))
	os.Setenv("INT64", fmt.Sprintf("%d", Int64))
	os.Setenv("UINT", fmt.Sprintf("%d", UInt))
	os.Setenv("UINT32", fmt.Sprintf("%d", UInt32))
	os.Setenv("UINT64", fmt.Sprintf("%d", UInt64))
	os.Setenv("FLOAT", fmt.Sprintf("%.2f", Float))
	os.Setenv("STRING", String)
	os.Setenv("SLICE", Slice)
	os.Setenv("SLICEINT", SliceInt)
	os.Setenv("SLICESTRING", SliceString)

	s := env.Source("", ",")

	bt, e := s.Bool("true")
	a.NoError(e)
	a.Equal(True, bt)

	bf, e := s.Bool("false")
	a.NoError(e)
	a.Equal(False, bf)

	i, e := s.Int("int")
	a.NoError(e)
	a.Equal(Int, i)

	i8, e := s.Int8("int8")
	a.NoError(e)
	a.Equal(Int8, i8)

	i32, e := s.Int32("int32")
	a.NoError(e)
	a.Equal(Int32, i32)

	i64, e := s.Int64("int64")
	a.NoError(e)
	a.Equal(Int64, i64)

	u, e := s.UInt("uint")
	a.NoError(e)
	a.Equal(UInt, u)

	u32, e := s.UInt32("uint32")
	a.NoError(e)
	a.Equal(UInt32, u32)

	u64, e := s.UInt64("uint64")
	a.NoError(e)
	a.Equal(UInt64, u64)

	f, e := s.Float("float")
	a.NoError(e)
	a.Equal(Float, f)

	str, e := s.String("string")
	a.NoError(e)
	a.Equal(String, str)

	sl, e := s.Slice("slice")
	a.NoError(e)
	a.Equal([]interface{}{"1", "a"}, sl)

	sli, e := s.SliceInt("sliceint")
	a.NoError(e)
	a.Equal([]int{1, 2, 3}, sli)

	slstr, e := s.SliceString("slicestring")
	a.NoError(e)
	a.Equal([]string{"a", "b", "c", "1", "2", "3"}, slstr)
}

func TestSource_WithPrefix(t *testing.T) {
	a := assert.New(t)
	var (
		True                = true
		False               = false
		Int         int     = 1
		Int8        int8    = 2
		Int32       int32   = 3
		Int64       int64   = 4
		UInt        uint    = 6
		UInt32      uint32  = 7
		UInt64      uint64  = 8
		Float       float64 = 3.14
		String              = "Foo bar"
		Slice               = "1,a"
		SliceInt            = "1,2,3"
		SliceString         = "a,b,c,1,2,3"
	)

	os.Setenv("GO_CFG_TRUE", "true")
	os.Setenv("GO_CFG_FALSE", "false")
	os.Setenv("GO_CFG_INT", fmt.Sprintf("%d", Int))
	os.Setenv("GO_CFG_INT8", fmt.Sprintf("%d", Int8))
	os.Setenv("GO_CFG_INT32", fmt.Sprintf("%d", Int32))
	os.Setenv("GO_CFG_INT64", fmt.Sprintf("%d", Int64))
	os.Setenv("GO_CFG_UINT", fmt.Sprintf("%d", UInt))
	os.Setenv("GO_CFG_UINT32", fmt.Sprintf("%d", UInt32))
	os.Setenv("GO_CFG_UINT64", fmt.Sprintf("%d", UInt64))
	os.Setenv("GO_CFG_FLOAT", fmt.Sprintf("%.2f", Float))
	os.Setenv("GO_CFG_STRING", String)
	os.Setenv("GO_CFG_SLICE", Slice)
	os.Setenv("GO_CFG_SLICEINT", SliceInt)
	os.Setenv("GO_CFG_SLICESTRING", SliceString)

	s := env.Source("GO_CFG", ",")

	bt, e := s.Bool("true")
	a.NoError(e)
	a.Equal(True, bt)

	bf, e := s.Bool("false")
	a.NoError(e)
	a.Equal(False, bf)

	i, e := s.Int("int")
	a.NoError(e)
	a.Equal(Int, i)

	i8, e := s.Int8("int8")
	a.NoError(e)
	a.Equal(Int8, i8)

	i32, e := s.Int32("int32")
	a.NoError(e)
	a.Equal(Int32, i32)

	i64, e := s.Int64("int64")
	a.NoError(e)
	a.Equal(Int64, i64)

	u, e := s.UInt("uint")
	a.NoError(e)
	a.Equal(UInt, u)

	u32, e := s.UInt32("uint32")
	a.NoError(e)
	a.Equal(UInt32, u32)

	u64, e := s.UInt64("uint64")
	a.NoError(e)
	a.Equal(UInt64, u64)

	f, e := s.Float("float")
	a.NoError(e)
	a.Equal(Float, f)

	str, e := s.String("string")
	a.NoError(e)
	a.Equal(String, str)

	sl, e := s.Slice("slice")
	a.NoError(e)
	a.Equal([]interface{}{"1", "a"}, sl)

	sli, e := s.SliceInt("sliceint")
	a.NoError(e)
	a.Equal([]int{1, 2, 3}, sli)

	slstr, e := s.SliceString("slicestring")
	a.NoError(e)
	a.Equal([]string{"a", "b", "c", "1", "2", "3"}, slstr)
}
