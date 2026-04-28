package detectors

import (
	"fmt"
	"sniping/adapters"
	"sniping/types"
)

func BasicDetector(resp types.ResponseData, input string, base types.ResponseData) {
	diff := false

	if resp.Status != base.Status {
		fmt.Printf("[diff] %s -> status=%d (base %d)\n",
			input, resp.Status, base.Status)
		adapters.Save(input)
		diff = true
	}

	if resp.Length != base.Length {
		fmt.Printf("[diff] %s -> length=%d (base %d)\n",
			input, resp.Length, base.Length)
		diff = true
	}

	if !diff {
		fmt.Printf("[fuzz] %s -> status=%d length=%d\n",
			input, resp.Status, resp.Length)
	}
}