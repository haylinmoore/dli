package providers

import (
	"fmt"
	"os"

	"github.com/libdns/loopia"
)

func getLoopiaProvider() (Provider, error) {
	username := os.Getenv("LOOPIA_API_USER")
	password := os.Getenv("LOOPIA_API_PASSWORD")

	if username == "" {
		return nil, fmt.Errorf("LOOPIA_API_USER environment variable not set")
	}
	if password == "" {
		return nil, fmt.Errorf("LOOPIA_API_PASSWORD environment variable not set")
	}

	provider := &loopia.Provider{
		Username: username,
		Password: password,
	}

	return provider, nil
}
