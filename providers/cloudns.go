package providers

import (
	"fmt"
	"os"

	"github.com/libdns/cloudns"
)

func getCloudnsProvider() (Provider, error) {
	authID := os.Getenv("CLOUDNS_AUTH_ID")
	authPassword := os.Getenv("CLOUDNS_AUTH_PASSWORD")
	subAuthID := os.Getenv("CLOUDNS_SUB_AUTH_ID")

	if authID == "" {
		return nil, fmt.Errorf("CLOUDNS_AUTH_ID environment variable not set")
	}
	if authPassword == "" {
		return nil, fmt.Errorf("CLOUDNS_AUTH_PASSWORD environment variable not set")
	}

	provider := &cloudns.Provider{
		AuthId:       authID,
		AuthPassword: authPassword,
		SubAuthId:    subAuthID, // Optional
	}

	return provider, nil
}
