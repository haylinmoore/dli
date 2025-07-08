package providers

import (
	"fmt"
	"os"

	"github.com/libdns/bunny"
)

func getBunnyProvider() (Provider, error) {
	accessKey := os.Getenv("BUNNY_API_KEY")
	if accessKey == "" {
		return nil, fmt.Errorf("BUNNY_API_KEY environment variable not set")
	}

	provider := &bunny.Provider{
		AccessKey: accessKey,
	}

	return provider, nil
}
