package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
	Type string
	Path string
	*viper.Viper
}

func (c *Config) Init()  {
	viper.SetConfigName(c.Name)
	viper.SetConfigType(c.Type)
	viper.AddConfigPath(c.Path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("ReadConfig err, ", err)
	}
}

func (c *Config) GetString(s string) string {
	return viper.GetString(s)
}
