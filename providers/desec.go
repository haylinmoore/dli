package providers

import (
	"fmt"
	"os"

	"github.com/libdns/desec"
)

func getDesecProvider() (Provider, error) {
	token := os.Getenv("DESEC_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DESEC_TOKEN environment variable not set")
	}

	provider := &desec.Provider{
		Token: token,
	}

	return provider, nil
}
