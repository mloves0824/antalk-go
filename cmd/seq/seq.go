package main

import (
	"antalk-go/internal/common"
	"antalk-go/internal/seq/protocol/http"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	configName = flag.String("config_name", "seq", "config name")
	configType = flag.String("config_type", "toml", "config type")
	configPath = flag.String("config_path", ".", "config path")
)

type server struct {
	name string
}

func main() {
	flag.Parse()

	c := &common.Config{
		Name: *configName,
		Type: *configType,
		Path: *configPath,
	}
	c.Init()

	httpSrv := http.New(c)

	// signal
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sg
		log.Printf("seq server get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			httpSrv.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
