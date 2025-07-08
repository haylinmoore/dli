package providers

import (
	"fmt"
	"os"

	"github.com/libdns/mijnhost"
)

func getMijnhostProvider() (Provider, error) {
	apiKey := os.Getenv("MIJNHOST_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("MIJNHOST_API_KEY environment variable not set")
	}

	provider := &mijnhost.Provider{
		ApiKey: apiKey,
	}

	return provider, nil
}
