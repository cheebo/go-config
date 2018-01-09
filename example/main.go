package main

import (
	cfg "github.com/cheebo/go-config"
	"github.com/cheebo/go-config/types"
	"github.com/davecgh/go-spew/spew"
)

type Config struct {
	Redis *types.RedisConfig

	Name string `description:"user's name'"`
	Pass string `description:"user's password'"`

	GasPeerTx float64

	Timeout        uint
	PricePerAction int

	AllowRegistration bool
}

func main() {
	c := Config{}
	cfgr := cfg.New()
	cfgr.Use(cfg.EnvironmentSource())
	cfgr.Use(cfg.FlagSource())
	cfgr.Configure(&c)
	spew.Dump(c)
}
