package main

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/env"
	"github.com/cheebo/go-config/sources/file"
	"github.com/cheebo/go-config/types"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	cfg := go_config.New()
	fs, err := file.Source(file.File{"./fixtures/config.json", go_config.JSON, ""})
	if err != nil {
		panic(err)
	}
	cfg.UseSource(env.Source("GO"), env.Source(""), fs)

	//fmt.Println(cfg.Get("name"), cfg.IsSet("name"))
	//fmt.Println(cfg.Get("age"), cfg.IsSet("age"))
	//fmt.Println(cfg.Get("amqp.url"), cfg.IsSet("amqp.url"))
	//fmt.Println(cfg.Get("amqp.addr"), cfg.IsSet("amqp.addr"))
	m := types.AMQPConfig{}
	err = cfg.Unmarshal(&m, "amqp")
	if err != nil {
		println(err.Error())
	}
	spew.Dump(m)
}
