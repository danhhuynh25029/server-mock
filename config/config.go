package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type hTTPConfig struct {
	Host   string `yaml:"host" env:"HTTP_HOST"`
	Port   int    `yaml:"port" env:"HTTP_PORT"`
	Origin string `yaml:"origin" env:"HTTP_ORIGIN"`
}

type mongoDBConfig struct {
	Host     string `json:"host" yaml:"host" env:"MONGO_HOST"`
	Port     int    `json:"port" yaml:"port" env:"MONGO_PORT"`
	User     string `json:"user" yaml:"user" env:"MONGO_USER"`
	Password string `json:"password" yaml:"password" env:"MONGO_PASSWORD"`
	Name     string `json:"name" yaml:"name" env:"MONGO_NAME"`
}

type allConfig struct {
	HTTP  hTTPConfig    `yaml:"http"`
	Mysql mongoDBConfig `yaml:"mongo"`
}

var Cfg allConfig

func Load() error {
	configFile := "config.yaml"
	if file := os.Getenv("CONFIG_FILE"); file != "" {
		configFile = file
	}
	err := cleanenv.ReadConfig(configFile, &Cfg)
	if err != nil {
		return err
	}
	return nil
}
