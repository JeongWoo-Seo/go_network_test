package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	return c
}
