package providers

import (
	"fmt"
	"os"

	"github.com/libdns/vultr/v2"
)

func getVultrProvider() (Provider, error) {
	apiToken := os.Getenv("VULTR_API_KEY")

	if apiToken == "" {
		return nil, fmt.Errorf("VULTR_API_KEY environment variable not set")
	}

	provider := &vultr.Provider{
		APIToken: apiToken,
	}

	return provider, nil
}
