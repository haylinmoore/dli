package providers

import (
	"fmt"
	"os"

	"github.com/libdns/ovh"
)

func getOvhProvider() (Provider, error) {
	applicationKey := os.Getenv("OVH_APPLICATION_KEY")
	applicationSecret := os.Getenv("OVH_APPLICATION_SECRET")
	consumerKey := os.Getenv("OVH_CONSUMER_KEY")
	endpoint := os.Getenv("OVH_ENDPOINT")

	if applicationKey == "" {
		return nil, fmt.Errorf("OVH_APPLICATION_KEY environment variable not set")
	}
	if applicationSecret == "" {
		return nil, fmt.Errorf("OVH_APPLICATION_SECRET environment variable not set")
	}
	if consumerKey == "" {
		return nil, fmt.Errorf("OVH_CONSUMER_KEY environment variable not set")
	}
	if endpoint == "" {
		return nil, fmt.Errorf("OVH_ENDPOINT environment variable not set")
	}

	provider := &ovh.Provider{
		ApplicationKey:    applicationKey,
		ApplicationSecret: applicationSecret,
		ConsumerKey:       consumerKey,
		Endpoint:          endpoint,
	}

	return provider, nil
}
