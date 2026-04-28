package app

import (
	"sniping/core"
)

func Run() {
	inputs := core.LoadInputs()
	cfg := core.InitConfig()
	
	core.RunEngine(inputs, cfg)
}