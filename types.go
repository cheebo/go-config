package go_config

type (
	Config interface {
		Fields

		UseSource(sources ...Source)
		SetDefault(key string, val interface{})
	}

	Fields interface {
		Get(key string) interface{}
		IsSet(key string) bool
		// sub fields
		Sub(key string) Fields
		// fields
		Bool(key string) bool
		Float(key string) float64
		Int(key string) int
		Int8(key string) int8
		Int32(key string) int32
		Int64(key string) int64
		Slice(key string) []interface{}
		SliceInt(key string) []int
		SliceString(key string) []string
		String(key string) string
		StringMap(key string) map[string]interface{}
		StringMapInt(key string) map[string]int
		StringMapSliceString(key string) map[string][]string
		StringMapString(key string) map[string]string
		UInt(key string) uint
		UInt32(key string) uint32
		UInt64(key string) uint64
		// unmarshal
		Unmarshal(v interface{}, prefix string) error
	}

	Source interface {
		Get(key string) interface{}
		IsSet(key string) bool
		// fields
		Bool(key string) (bool, error)
		Float(key string) (float64, error)
		Int(key string) (int, error)
		Int8(key string) (int8, error)
		Int32(key string) (int32, error)
		Int64(key string) (int64, error)
		SliceString(key string) ([]string, error)
		SliceInt(key string) ([]int, error)
		String(key string) (string, error)
		StringMap(key string) map[string]interface{}
		StringMapInt(key string) map[string]int
		StringMapSliceString(key string) map[string][]string
		StringMapString(key string) map[string]string
		UInt(key string) (uint, error)
		UInt32(key string) (uint32, error)
		UInt64(key string) (uint64, error)
	}
)
