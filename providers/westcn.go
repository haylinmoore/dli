package providers

import (
	"fmt"
	"os"

	"github.com/libdns/westcn"
)

func getWestcnProvider() (Provider, error) {
	username := os.Getenv("WESTCN_USERNAME")
	password := os.Getenv("WESTCN_API_PASSWORD")

	if username == "" {
		return nil, fmt.Errorf("WESTCN_USERNAME environment variable not set")
	}
	if password == "" {
		return nil, fmt.Errorf("WESTCN_API_PASSWORD environment variable not set")
	}

	provider := &westcn.Provider{
		Username:    username,
		APIPassword: password,
	}

	return provider, nil
}
