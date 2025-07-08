package providers

import (
	"fmt"
	"os"

	"github.com/libdns/hetzner"
)

func getHetznerProvider() (Provider, error) {
	apiKey := os.Getenv("HETZNER_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("HETZNER_API_KEY environment variable not set")
	}

	provider := &hetzner.Provider{
		AuthAPIToken: apiKey,
	}

	return provider, nil
}
