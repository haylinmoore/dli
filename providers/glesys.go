package providers

import (
	"fmt"
	"os"

	"github.com/libdns/glesys"
)

func getGlesysProvider() (Provider, error) {
	apiUser := os.Getenv("GLESYS_API_USER")
	apiKey := os.Getenv("GLESYS_API_KEY")

	if apiUser == "" {
		return nil, fmt.Errorf("GLESYS_API_USER environment variable not set")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("GLESYS_API_KEY environment variable not set")
	}

	provider := &glesys.Provider{
		Project: apiUser,
		APIKey:  apiKey,
	}

	return provider, nil
}
