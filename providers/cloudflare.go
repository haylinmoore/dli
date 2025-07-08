package providers

import (
	"fmt"
	"os"

	"github.com/libdns/cloudflare"
)

func getCloudflareProvider() (Provider, error) {
	// Try primary environment variables first
	apiToken := os.Getenv("CF_DNS_API_TOKEN")
	if apiToken == "" {
		// Try alias variables
		apiToken = os.Getenv("CLOUDFLARE_DNS_API_TOKEN")
	}

	zoneToken := os.Getenv("CF_ZONE_API_TOKEN")
	if zoneToken == "" {
		zoneToken = os.Getenv("CLOUDFLARE_ZONE_API_TOKEN")
	}

	if apiToken == "" {
		return nil, fmt.Errorf("CF_DNS_API_TOKEN or CLOUDFLARE_DNS_API_TOKEN environment variable not set")
	}

	provider := &cloudflare.Provider{
		APIToken:  apiToken,
		ZoneToken: zoneToken, // Optional
	}

	return provider, nil
}
