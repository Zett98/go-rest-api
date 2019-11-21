package apiserver

import "go-rest-api/internal/app/store"

// Config ... server config
type Config struct {
	bindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ... returns config for server
func NewConfig() *Config {
	return &Config{
		bindAddr: ":8080",
		logLevel: "debug",
		Store:    store.NewConfig(),
	}
}
