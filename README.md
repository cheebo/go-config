# Go Config

Package provdes routines that loads configuration into provided structure from provided sources.


Example

```go
type Config struct {
	Redis *types.RedisConfig

	Name  string `description:"user's name'"`
	Pass  string `description:"user's password'"`

	GasPeerTx           float64

	Timeout             uint
	PricePerAction      int

	AllowRegistration   bool
}

func main() {
	c := Config{}
	cfgr := configurer.New()
	// parse env variables
	cfgr.Use(configurer.EnvironmentSource())
	// parse flags
	cfgr.Use(configurer.FlagSource())
	cfgr.Configure(&c)
}
```