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
	IsProduction bool   `yaml:"isProduction"` //Template reloads, serve adds/analytics
	Port         string `yaml:"port"`         //`env:"PORT" env-default:"9000"`
	Domain       string `yaml:"domain"`       //`env:"DOMAIN" env-default:"http://127.0.0.1:9000"`
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
