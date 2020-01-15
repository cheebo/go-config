package main

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/internal/reader"
	"github.com/cheebo/go-config/sources/env"
	"github.com/cheebo/go-config/sources/file"
)

type AMQPConfig struct {
	URL          string `json:"url"`
	Exchange     string `json:"exchange"`
	Queue        string `json:"queue"`
	Kind         string `json:"kind"`
	Key          string `json:"key"`
	Durable      bool   `json:"durable"`
	AutoDelete   bool   `json:"auto_delete"`
	DeliveryMode uint   `json:"delivery_mode"`
}

func main() {
	cfg := go_config.New()
	fileSource, err := file.Source(file.File{"./fixtures/config.json", reader.JSON, ""})
	if err != nil {
		panic(err)
	}
	cfg.UseSource(env.Source("GO"), env.Source(""), fileSource)

	amqp := AMQPConfig{}
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
