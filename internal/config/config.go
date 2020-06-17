package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PathPrefix   string `default:"" split_words:"true"`
	RedisURL     string `default:"redis://localhost" split_words:"true"`
	RedisCluster bool   `default:"false" split_words:"true"`
}

func FromEnv() (Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	return c, err
}
