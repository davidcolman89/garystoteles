package config

import (
	"github.com/davidcolman89/config"
)

type Config struct {
	ServerAddr string
	Host string
	Port string
	Network string
}

func NewConfig(prefix string) Config {
	return initVipper(prefix)
}

func initVipper(prefix string) Config {
	c := Config{}
	config.AddConfigPath("/etc/config")
	config.SetConfig(prefix, &c, "config")
	return c
}


