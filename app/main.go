package main

import (
	"cleans"
	"config"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var (
	conf  *config.Config
	DEBUG bool
)

func init() {
	var confPath string
	flag.StringVar(&confPath, "config-path", "cleanDirs.toml", "Path to config file")
	flag.BoolVar(&DEBUG, "debug", false, "Enable debug mode")
	flag.Parse()

	DEBUG = true

	fmt.Println(DEBUG)
	conf = config.NewConfig()
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		fmt.Println("Концифиг файл не найден: ", err.Error())
		os.Exit(1)
	}
}

func main() {
	//Delete Scans
	err := cleans.CleanScans(conf.Scans.NameF, conf.Formats.NameFFormats, conf.Files.Excepts, DEBUG)
	if err != nil {
		fmt.Println("Error cleaning scans:", err.Error())
	}
	fmt.Println()
	err = cleans.CleanUsers(conf.Users.NameF, conf.Formats.NameFFormats, DEBUG)
	if err != nil {
		fmt.Println("Error cleaning users:", err.Error())
	}
}
