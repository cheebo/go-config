# What is GoConfig

GoConfig is a complete configuration solution for Go applications.

It supports:
- setting default values
- reading from environment variables
- reading from remote config systems (Consul)
- can parse JSON, TOML, YAML configurations
- extendable configuration sources: you can write your own config loader for unsupported file format, remote etc.
- can define prefix for variables from specific configuration source

### Config Sources
GoConfig knows nothing about environment variables, config files or remote configuration systems.
It's empty by default. 
It can load variables from provided Config Sources. 
Config Source could be env variables, flags, files, remote etc.
GoConfig provides method `UseSource` to attach needed Config Source and
it has no attached config sources by default.

GoConfig provides implementations for several Config Sources to load variables from environment, from JSON, 
TOML, YAML files and Consul. 
They could be found in `sources` directory.

If you want to get access to environment variables, then you must attach env config source:
```go

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/env"
)

func main() {
	cfg := go_config.New()
	cfg.UseSource(env.Source("MYAPP"))
	// now you can get value of env variables
	// get MYAPP_HTTP_HOST as string
	cfg.String("http.host")
}
```

If you want to get access to configuration from specific file, then you must attach file config source:

```go

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/file"
)

func main() {
    cfg := go_config.New()
    fs, err := file.Source(
        file.File{"./config.json", go_config.JSON, ""},
    )
    if err != nil {
        panic(err)
    }
    // now you can get value from file 
    cfg.String("http_host")
}
```

### Loading Order and overriding

GoConfig provides `UseSource` method to use values from specific config sources (env, files, remote etc).
The Order of sources provided as function parameters is the order in which sources are processed.
If two sources have a variable with the same name or path, then the value of the variable from the last source is taken.

For example, we have two config sources:
```go
cfg := go_config.New()

// create Source to load values from environment variables
envs := env.Source("GO")

// create Source to load values from file
fs, err := file.Source(
    file.File{"./config.json", go_config.JSON, ""},
)
if err != nil {
    panic(err)
}
```

Now we can use this sources, but attachment order matters. 

If we attach env source first, and then attach file source, 
```go
cfg.UseSource(envs, fs)
```
variables from file will override variables from env if 
they have equal names or paths.

If we attach file first, and then attach env source, 
```go
cfg.UseSource(fs, env)
```
env variables will override variables from file if 
they have equal names or paths. 

### Config source namespaces

@TODO

### Getting values
There are several methods exist to get a value depending on the value's type:

- Bool(key string) bool
- Get(key string) interface{}
- Float(key string) float64
- Int(key string) int
- IsSet(key string) bool
- Slice(key, delimiter string) []interface{}
- String(key string) string
- StringMap(key string) map[string]interface{}
- Sub(key string) Fields
- UInt(key string) uint

### Variable names and nested keys
You can use any readable string as a variable name: "name", "http_port" or "core_variable_name#1".

GoConfig methods accept formatted paths to nested keys.
For example, if the following JSON file is loaded:
```json
{
  "http": {
    "host": "localhost",
    "port": 8080
  }
}
```

GoConfig can access a nested field by passing a `.` delimited path of keys:

`cfg.String("http.host")` or `cfg.Int("http.port")`


### Extract sub-tree
For example, GoConfig represents file
```json
{
  "server": {
      "http": {
        "host": "127.0.0.1",
        "port": 8080
      },
      "grpc": {
        "host": "0.0.0.0",
        "port": 8081
    }
  }
}
```

after executing
```go
srv := cfg.Sub("server.http")
```

`srv` represents

```json
{
  "host": "127.0.0.1",
  "port": 8080
}
```
