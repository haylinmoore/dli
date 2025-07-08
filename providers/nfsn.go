package providers

import (
	"fmt"
	"os"

	"github.com/libdns/nfsn"
)

func getNfsnProvider() (Provider, error) {
	login := os.Getenv("NEARLYFREESPEECH_LOGIN")
	apiKey := os.Getenv("NEARLYFREESPEECH_API_KEY")

	if login == "" {
		return nil, fmt.Errorf("NEARLYFREESPEECH_LOGIN environment variable not set")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("NEARLYFREESPEECH_API_KEY environment variable not set")
	}

	provider := &nfsn.Provider{
		Login:  login,
		APIKey: apiKey,
	}

	return provider, nil
}
