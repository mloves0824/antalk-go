package main

import (
	"antalk-go/internal/common"
	"antalk-go/internal/push/controller/rpcx"
	"flag"
	"log"
)

var (
	configName = flag.String("config_name", "push", "config name")
	configType = flag.String("config_type", "toml", "config type")
	configPath = flag.String("config_path", ".", "config path")
)

func main() {
	flag.Parse()
	c := &common.Config{
		Name: *configName,
		Type: *configType,
		Path: *configPath,
	}
	c.Init()

	_, err := rpcx.New(c)
	if err != nil {
		log.Fatal("rpcx.New error, ", err)
	}
}
