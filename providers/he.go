package providers

import (
	"fmt"
	"os"

	"github.com/libdns/he"
)

func getHeProvider() (Provider, error) {
	// Hurricane Electric uses different variable name in lego vs libdns
	apiKey := os.Getenv("HURRICANE_TOKENS")
	if apiKey == "" {
		// Try alternative naming
		apiKey = os.Getenv("HE_API_KEY")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("HURRICANE_TOKENS or HE_API_KEY environment variable not set")
	}

	provider := &he.Provider{
		APIKey: apiKey,
	}

	return provider, nil
}
