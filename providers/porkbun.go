package providers

import (
	"fmt"
	"os"

	"github.com/libdns/porkbun"
)

func getPorkbunProvider() (Provider, error) {
	apiKey := os.Getenv("PORKBUN_API_KEY")
	secretKey := os.Getenv("PORKBUN_SECRET_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("PORKBUN_API_KEY environment variable not set")
	}
	if secretKey == "" {
		return nil, fmt.Errorf("PORKBUN_SECRET_API_KEY environment variable not set")
	}

	provider := &porkbun.Provider{
		APIKey:       apiKey,
		APISecretKey: secretKey,
	}

	return provider, nil
}
