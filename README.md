# Go Config

Package provides routines that loads configuration into provided structure from provided sources.


Example

```go
type Config struct {
	Redis *types.RedisConfig

	Name  string `description:"user's name'"`
	Pass  string `cfg:"password" description:"user's password'"`

	GasPeerTx           float64

	Timeout             uint
	PricePerAction      int

	AllowRegistration   bool
}

func main() {
	cfg := go_config.New()
	
	// use config file
    fs, err := file.Source(
    	file.File{"./fixtures/config.json", go_config.JSON, ""},
    )
    if err != nil {
        panic(err)
    }
    
    // use environment variables and file config
    // the order is important: later sources override values from previous sources
    cfg.UseSource(env.Source("GO"), env.Source(""), fs)
    
    // get variables and isSet state
    fmt.Println(cfg.Get("name"), cfg.IsSet("name"))
    fmt.Println(cfg.Get("amqp.url"), cfg.IsSet("amqp.url"))
    fmt.Println(cfg.Get("amqp.addr"), cfg.IsSet("amqp.addr"))
    
    // Get ENV variable
    fmt.Println(cfg.Get("home"), cfg.IsSet("home"))
    
    // Unmarshal config data into structure
    m := types.AMQPConfig{}
    err = cfg.Unmarshal(&m, "amqp")
    if err != nil {
        println(err.Error())
    }
}
```

## Config Source

[futured] Config source (cs) is the flag that defines configuration source.

```bash
./service -cs="cs=<type>,opt=arg,opt[=arg];<type>,opt=arg,..."
```

Supported config sources:
- environment variables
- flags (FIXME: read data from flags)
- file
  - json
  - yaml
  - toml
- consul