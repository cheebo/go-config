package main

import (
	"fmt"
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/sources/env"
	"github.com/cheebo/go-config/sources/file"
)

func main() {
	cfg := go_config.New()
	fs, err := file.Source("./fixtures/config.json", go_config.JSON)
	if err != nil {
		panic(err)
	}
	cfg.UseSource(env.Source("GO"), env.Source(""), fs)
	fmt.Println(cfg.Get("name"), cfg.IsSet("name"))
	fmt.Println(cfg.Get("amqp.url"), cfg.IsSet("amqp.url"))
	fmt.Println(cfg.String("amqp.url2"), cfg.IsSet("amqp.url2"))
	fmt.Println(cfg.Int("amqp.port"), cfg.IsSet("amqp.port"))
	fmt.Println(cfg.Get("home"), cfg.IsSet("home"), cfg.IsSet("myhome"))
}
