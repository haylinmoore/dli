package providers

import (
	"fmt"
	"os"

	"github.com/libdns/digitalocean"
)

func getDigitaloceanProvider() (Provider, error) {
	apiToken := os.Getenv("DO_AUTH_TOKEN")

	if apiToken == "" {
		return nil, fmt.Errorf("DO_AUTH_TOKEN environment variable not set")
	}

	provider := &digitalocean.Provider{
		APIToken: apiToken,
	}

	return provider, nil
}
