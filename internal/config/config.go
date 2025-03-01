package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HostConfig     `yaml:"host"`
	BotConfig      `yaml:"bot"`
	DatabaseConfig `yaml:"database"`
}

type BotConfig struct {
	Key string `yaml:"key"`
}
type HostConfig struct {
	IsProduction bool   `yaml:"isProduction"`
	Port         string `yaml:"port"`
	Domain       string `yaml:"domain"`
	Title        string `yaml:"title"`
}
type DatabaseConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Name           string `yaml:"name"`
	MaxConnections int32  `yaml:"maxConnections"`
	Sslmode        string `yaml:"sslmode"`
}

func New(path string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
