package providers

import (
	"fmt"
	"os"

	"github.com/libdns/netcup"
)

func getNetcupProvider() (Provider, error) {
	customerNumber := os.Getenv("NETCUP_CUSTOMER_NUMBER")
	apiKey := os.Getenv("NETCUP_API_KEY")
	apiPassword := os.Getenv("NETCUP_API_PASSWORD")

	if customerNumber == "" {
		return nil, fmt.Errorf("NETCUP_CUSTOMER_NUMBER environment variable not set")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("NETCUP_API_KEY environment variable not set")
	}
	if apiPassword == "" {
		return nil, fmt.Errorf("NETCUP_API_PASSWORD environment variable not set")
	}

	provider := &netcup.Provider{
		CustomerNumber: customerNumber,
		APIKey:         apiKey,
		APIPassword:    apiPassword,
	}

	return provider, nil
}
