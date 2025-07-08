package providers

import (
	"fmt"
	"os"

	"github.com/libdns/dynu"
)

func getDynuProvider() (Provider, error) {
	apiKey := os.Getenv("DYNU_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("DYNU_API_KEY environment variable is required")
	}

	provider := &dynu.Provider{
		APIToken: apiKey,
	}

	return provider, nil
}
