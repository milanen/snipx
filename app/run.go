package app

import (
	"sniping/core"
)

func Run(configPath string) {
	cfg := core.InitConfig(configPath)
	inputs := core.LoadInputs(cfg)
	
	core.RunEngine(inputs, cfg)
}