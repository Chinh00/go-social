package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go-social/internal/pkg/core"
	"os"
)

type Config struct {
	core.App
	core.Http
}

func NewConfig() (*Config, error) {
	config := &Config{}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	err = cleanenv.ReadConfig(dir+"/config.yml", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
