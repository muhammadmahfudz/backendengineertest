package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App        `yaml:"app"`
		PostgreSQL `yaml:"postgresql"`
		Http       `yaml:"http"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	PostgreSQL struct {
		Host   string `env-required:"true" yaml:"host" env:"HOST"`
		Port   int    `env-required:"true" yaml:"port" env:"PORT"`
		User   string `env-required:"true" yaml:"dbuser" env:"DBUSER"`
		Pass   string `env-required:"true" yaml:"pass" env:"PASS"`
		DbName string `env-required:"true" yaml:"dbname" env:"DBNAME"`
	}

	Http struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}
)

func NewConfig() (*Config, error) {

	cfg := &Config{}
	path := "./config/local.yaml"

	err := cleanenv.ReadConfig(path, cfg)

	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}

	fmt.Println(cfg.App.Version)

	return cfg, nil
}
