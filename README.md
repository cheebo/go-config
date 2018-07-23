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
    fs, err := file.Source("./fixtures/config.json", go_config.JSON)
    if err != nil {
        panic(err)
    }
    // use environment variables and file config
    cfg.UseSource(env.Source("GO"), env.Source(""), fs)
    // get variables and isSet state
    fmt.Println(cfg.Get("name"), cfg.IsSet("name"))
    fmt.Println(cfg.Get("amqp.url"), cfg.IsSet("amqp.url"))
    fmt.Println(cfg.Get("amqp.url2"), cfg.IsSet("amqp.url2"))
    fmt.Println(cfg.Get("home"), cfg.IsSet("home"), cfg.IsSet("myhome"))
}
```

## Config Source

Config source (cs) is the flag that defines configuration source.

```bash
./service -cs="cs=<type>,opt=arg,opt[=arg];<type>,opt=arg,..."
```

Supported config sources:
- environment variables
- flags (FIXME: read flags)
- file
  - json
  - yaml
  - toml
- consul (FIXME: implement data import from consul)

#### Consul CS

```
cs="consul;"
```

## Tags


Supported field tags:
- cfg:"param_name"
- description:"variable description"
- default:"default_value"
