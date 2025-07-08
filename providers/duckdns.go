package providers

import (
	"fmt"
	"os"

	"github.com/libdns/duckdns"
)

func getDuckdnsProvider() (Provider, error) {
	token := os.Getenv("DUCKDNS_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DUCKDNS_TOKEN environment variable is required")
	}

	provider := &duckdns.Provider{
		APIToken: token,
	}

	return provider, nil
}
