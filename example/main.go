package main

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/env"
	"github.com/cheebo/go-config/sources/file"
	"github.com/cheebo/go-config/types"
)

func main() {
	cfg := go_config.New()
	fileSource, err := file.Source(file.File{"./fixtures/config.json", go_config.JSON, ""})
	if err != nil {
		panic(err)
	}
	cfg.UseSource(env.Source("GO"), env.Source(""), fileSource)

	amqp := types.AMQPConfig{}
	err = cfg.Unmarshal(&amqp, "amqp")
	if err != nil {
		println(err.Error())
	}

	sub := cfg.Sub("amqp")
	println(sub.String("url"))

	println(cfg.String("GOPATH"))

	sub = cfg.Sub("tree.root")
	s := sub.Sub("branch")
	println(s.String("key"))
}
