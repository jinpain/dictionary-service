package config

import "github.com/ilyakaznacheev/cleanenv"

type Server struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
}

func LoadServer(path string) *Server {
	var cfg Server

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read db config: " + err.Error())
	}

	return &cfg
}
