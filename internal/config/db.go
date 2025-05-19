package config

import "github.com/ilyakaznacheev/cleanenv"

type DB struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DbName   string `yaml:"db_name" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-required:"true"`
}

func LoadDB(path string) *DB {
	var cfg DB

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read db config: " + err.Error())
	}

	return &cfg
}
