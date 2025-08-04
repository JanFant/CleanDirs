package config

type Users struct {
	NameF       []string `toml:"UsersName"`
	DirsExcepts []string `toml:"DirsExcept"`
}
