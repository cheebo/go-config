package main

import (
	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/pkg/sources/env"
	"github.com/cheebo/go-config/pkg/sources/file"
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
	fileSource, err := file.Source(file.File{"/Users/sergey/Documents/development/sources/cheebo/go-config/test/testdata/config.json", ""})
	if err != nil {
		panic(err)
	}
	cfg.UseSource(env.Source("GO", ","), env.Source("", ","), fileSource)

	amqp := AMQPConfig{}
	err = cfg.Unmarshal(&amqp, "amqp")
	if err != nil {
		println(err.Error())
	}
	println(amqp.URL, amqp.Kind)

	sub := cfg.Sub("amqp")
	println(sub.String("url"), sub.String("kind"))

	sub = cfg.Sub("tree.root")
	s := sub.Sub("branch")
	println("tree.root.branch: ", s.String("key"))
}
