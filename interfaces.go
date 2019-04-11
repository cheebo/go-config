package go_config

type Config interface {
	Fields

	UseSource(sources ...Source)
	SetDefault(key string, val interface{})
}

type Fields interface {
	Bool(key string) bool
	Get(key string) interface{}
	Float(key string) float64
	Int(key string) int
	IsSet(key string) bool
	Slice(key, delimiter string) []interface{}
	String(key string) string
	StringMap(key string) map[string]interface{}
	Sub(key string) Fields
	UInt(key string) uint
	Unmarshal(v interface{}, prefix string) error
}

type Source interface {
	Get(key string) interface{}

	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	Int(key string) (int, error)
	UInt(key string) (uint, error)
	Slice(key, delimiter string) ([]interface{}, error)
	String(key string) (string, error)
	StringMap(key string) map[string]interface{}

	IsSet(key string) bool
}
