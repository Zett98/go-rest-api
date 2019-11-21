package store

// Config ... DB config
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig ... returns new db config
func NewConfig() *Config {
	return &Config{}
}
