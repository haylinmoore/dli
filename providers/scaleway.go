package providers

import (
	"fmt"
	"os"

	"github.com/libdns/scaleway"
)

func getScalewayProvider() (Provider, error) {
	secretKey := os.Getenv("SCW_SECRET_KEY")
	projectID := os.Getenv("SCW_PROJECT_ID")

	if secretKey == "" {
		return nil, fmt.Errorf("SCW_SECRET_KEY environment variable not set")
	}
	if projectID == "" {
		return nil, fmt.Errorf("SCW_PROJECT_ID environment variable not set")
	}

	provider := &scaleway.Provider{
		SecretKey:      secretKey,
		OrganizationID: projectID,
	}

	return provider, nil
}
