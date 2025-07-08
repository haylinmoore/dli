package providers

import (
	"fmt"
	"os"

	"github.com/libdns/dnsimple"
)

func getDnsimpleProvider() (Provider, error) {
	apiToken := os.Getenv("DNSIMPLE_OAUTH_TOKEN")
	baseURL := os.Getenv("DNSIMPLE_BASE_URL")

	if apiToken == "" {
		return nil, fmt.Errorf("DNSIMPLE_OAUTH_TOKEN environment variable not set")
	}

	provider := &dnsimple.Provider{
		APIAccessToken: apiToken,
		APIURL:         baseURL, // Optional
	}

	return provider, nil
}
