package main

import (
	"flag"
	"sniping/app"
)

func main() {
	configPath := flag.String(
		"c",
		"config/models/config.yaml",
		"path to config file",
	)

	flag.Parse()

	app.Run(*configPath)
}