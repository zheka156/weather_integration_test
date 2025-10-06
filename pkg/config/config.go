package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigPath = "../configs/app.yaml"

type Config struct {
	WeatherClient WeatherClient `yaml:"weatherClient"`
}

type WeatherClient struct {
	URL string `yaml:"url"`
}

func LoadConfig() *Config {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
