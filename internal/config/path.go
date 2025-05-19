package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Path struct {
	DB     string `yaml:"db_path" env-required:"true"`
	Server string `yaml:"server_path" env-required:"true"`
}

func LoadPath() *Path {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Path

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
