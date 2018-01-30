package main

import (
	cfg "github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/davecgh/go-spew/spew"
)

type Master struct {
	AMQP types.AMQPConfig `consul:"amqp"`
}

type Config struct {
	Master Master

	Name string `description:"user's name'"`
	Pass string `cfg:"password" description:"user's password'"`

	GasPeerTx float64 `default:"10.11"`

	Timeout        uint `default:"101"`
	PricePerAction int

	AllowRegistration bool `default:"true"`

	Ips []string
}

func main() {
	c := Config{}
	cfgr := cfg.New()
	//cfgr.Use(cfg.EnvironmentSource())
	//cfgr.Use(cfg.ConsulSource("/example/config", types.ConsulConfig{
	//	Addr:   "localhost:8500",
	//	Scheme: "http",
	//}))
	cfgr.Use(cfg.FlagSource())
	cfgr.Configure(&c)

	spew.Dump(c)
}
