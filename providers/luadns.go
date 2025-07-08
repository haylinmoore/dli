package providers

import (
	"fmt"
	"os"

	"github.com/libdns/luadns"
)

func getLuadnsProvider() (Provider, error) {
	username := os.Getenv("LUADNS_API_USERNAME")
	apiToken := os.Getenv("LUADNS_API_TOKEN")

	if username == "" {
		return nil, fmt.Errorf("LUADNS_API_USERNAME environment variable not set")
	}
	if apiToken == "" {
		return nil, fmt.Errorf("LUADNS_API_TOKEN environment variable not set")
	}

	provider := &luadns.Provider{
		Email:  username,
		APIKey: apiToken,
	}

	return provider, nil
}
