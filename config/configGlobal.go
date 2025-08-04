package config

var GlobalConfig *Config

// Config struct all configs
type Config struct {
	Scans   Scans   `toml:"Scans"`
	Users   Users   `toml:"Users"`
	Formats Formats `toml:"Formats"`
	Files   Files   `toml:"Files"`
}

func NewConfig() *Config {
	return &Config{}
}
