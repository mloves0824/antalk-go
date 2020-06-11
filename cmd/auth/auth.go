package main

import (
	"antalk-go/internal/auth"
	"antalk-go/internal/auth/config"
	"flag"
	"fmt"
	"log"
)

var (
	configName = flag.String("config_name", "auth", "config name")
	configType = flag.String("config_type", "toml", "config type")
	configPath = flag.String("config_path", ".", "config path")
)

func init()  {
	fmt.Println("test")
}

func main() {
	flag.Parse()
	c := &config.Config{
		Name: *configName,
		Type: *configType,
		Path: *configPath,
	}
	c.Init()

	s, err := auth.New(c)
	if err != nil {
		log.Fatal("auth.New error, ", err)
	}
	defer s.Close()
}