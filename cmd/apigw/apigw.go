package main

import (
	"antalk-go/internal/apigw"
	"antalk-go/internal/common"
	"flag"
	"log"
)

var (
	configName = flag.String("config_name", "apigw", "config name")
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

	s, err := apigw.NewApigw(c)
	if err != nil {
		log.Fatal("NewApigw error, ", err)
	}
	defer s.Close()
}
