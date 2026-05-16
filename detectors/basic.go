package detectors

import (
	"fmt"
	"sniping/adapters"
	"sniping/config"
	"sniping/types"
)

func BasicDetector(resp types.ResponseData, input string, base types.ResponseData, cfg config.Detector) {
	diff := false

	// explicit matchStatus
	if len(cfg.MatchStatus) > 0 {
		for _, status := range cfg.MatchStatus {
			if resp.Status == status {
				fmt.Printf("[match] %s -> status=%d\n", input, resp.Status)
				adapters.Save(input)
				return
			}
		}

		return
	}

	// explicit excludeSize
	for _, size := range cfg.ExcludeSize {
		if resp.Length == size {
			return
		}
	}

	// baseline fallback
	if resp.Status != base.Status {
		fmt.Printf("[diff] %s -> status=%d (base %d)\n",
			input,
			resp.Status,
			base.Status,
		)
		adapters.Save(input)
		diff = true
	}

	if resp.Length != base.Length {
		fmt.Printf("[diff] %s -> length=%d (base %d)\n",
			input,
			resp.Length,
			base.Length,
		)
		diff = true
	}

	if !diff {
		fmt.Printf("[fuzz] %s -> status=%d length=%d\n",
			input,
			resp.Status,
			resp.Length,
		)
	}
}