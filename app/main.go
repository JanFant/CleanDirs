package main

import (
	"config"
	"flag"
	"fmt"
)

func init() {
	var confPath string
	flag.StringVar(&confPath, "config-path", "", "Path to config file")

	config.GlobalConfig = config.NewConfig()
}

func main() {
	fmt.Println("hell yeah")
	fmt.Println("Gi")
	fmt.Println("Test Test")
}
