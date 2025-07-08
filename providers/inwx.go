package providers

import (
	"fmt"
	"os"

	"github.com/libdns/inwx"
)

func getInwxProvider() (Provider, error) {
	username := os.Getenv("INWX_USERNAME")
	password := os.Getenv("INWX_PASSWORD")
	sharedSecret := os.Getenv("INWX_SHARED_SECRET")

	if username == "" {
		return nil, fmt.Errorf("INWX_USERNAME environment variable not set")
	}
	if password == "" {
		return nil, fmt.Errorf("INWX_PASSWORD environment variable not set")
	}

	provider := &inwx.Provider{
		Username:     username,
		Password:     password,
		SharedSecret: sharedSecret, // Optional
	}

	return provider, nil
}
