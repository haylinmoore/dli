package providers

import (
	"fmt"
	"os"

	"github.com/libdns/metaname"
)

func getMetanameProvider() (Provider, error) {
	apiKey := os.Getenv("METANAME_API_KEY")
	accountRef := os.Getenv("METANAME_ACCOUNT_REFERENCE")

	if apiKey == "" {
		return nil, fmt.Errorf("METANAME_API_KEY environment variable not set")
	}
	if accountRef == "" {
		return nil, fmt.Errorf("METANAME_ACCOUNT_REFERENCE environment variable not set")
	}

	provider := &metaname.Provider{
		APIKey:           apiKey,
		AccountReference: accountRef,
	}

	return provider, nil
}
