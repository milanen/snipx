package core

import (
	"fmt"
	"sniping/adapters"
	"sniping/config"
	"sniping/detectors"
	"strings"
)

func RunEngine(inputs []string, cfg config.Config) {
	client, err := CreateCustomClient(adapters.GetProxy(), cfg.Network.RequestTimeout, cfg.Network.EnableGateway)
		if err != nil {
			fmt.Printf("[client] err: %s\n", err)
			return
			}

	// baseline
	baseline := strings.ReplaceAll(cfg.Target.Template, "§", inputs[0])

	base, err := DoRequest(client, "GET", baseline, nil)
		if err != nil {
			fmt.Printf("[base] err: %s\n", err)
			return			
		}

	if base.Body == nil {
    	fmt.Printf("[base] err: %s\n", err)
	}

	fmt.Printf("[base] %s -> status=%d length=%d\n", inputs[0], base.Status, base.Length)

	//loop
	for _, el := range inputs {
		target := strings.ReplaceAll(cfg.Target.Template, "§", el)

		resp, err := DoRequest(client, "GET", target, nil)
		if err != nil {
			fmt.Printf("[fuzz] err: %s\n", err)
			return			
		}
		// run detector
		detectors.BasicDetector(resp, el, base)
	}


}