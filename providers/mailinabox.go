package providers

import (
	"fmt"
	"os"

	"github.com/libdns/mailinabox"
)

func getMailinaboxProvider() (Provider, error) {
	baseURL := os.Getenv("MAILINABOX_BASE_URL")
	email := os.Getenv("MAILINABOX_EMAIL")
	password := os.Getenv("MAILINABOX_PASSWORD")

	if baseURL == "" {
		return nil, fmt.Errorf("MAILINABOX_BASE_URL environment variable not set")
	}
	if email == "" {
		return nil, fmt.Errorf("MAILINABOX_EMAIL environment variable not set")
	}
	if password == "" {
		return nil, fmt.Errorf("MAILINABOX_PASSWORD environment variable not set")
	}

	provider := &mailinabox.Provider{
		APIURL:       baseURL,
		EmailAddress: email,
		Password:     password,
	}

	return provider, nil
}
