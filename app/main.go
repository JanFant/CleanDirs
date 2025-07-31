package main

import (
	"cleans"
	"config"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func init() {
	var confPath string
	flag.StringVar(&confPath, "config-path", "cleanDirs.toml", "Path to config file")

	config.GlobalConfig = config.NewConfig()
	if _, err := toml.DecodeFile(confPath, &config.GlobalConfig); err != nil {
		fmt.Println("Error loading config file:", err.Error())
		os.Exit(1)
	}
	fmt.Println(config.GlobalConfig)
}

func main() {
	//Delete Scans
	err := cleans.CleanScans(config.GlobalConfig.ScansConfig.NameFScans, config.GlobalConfig.FormatsConfig.NameFFormats)
	if err != nil {
		fmt.Println("Error cleaning scans:", err.Error())
	}
	err = cleans.CleanUsers(config.GlobalConfig.UsersConfig.NameFUsers, config.GlobalConfig.FormatsConfig.NameFFormats)
	if err != nil {
		fmt.Println("Error cleaning users:", err.Error())
	}
}
