package config

var GlobalConfig *Config

// Config struct all configs
type Config struct {
	ScansConfig ConfigScans `toml:"Scans"`
	UsersConfig ConfigUsers `toml:"Users"`
	DocsConfig  ConfigDocs  `toml:"Docs"`
}

func NewConfig() *Config {
	return &Config{}
}
