package providers

import (
	"fmt"
	"os"

	"github.com/libdns/googleclouddns"
)

func getGoogleclouddnsProvider() (Provider, error) {
	project := os.Getenv("GCE_PROJECT")
	serviceAccount := os.Getenv("GCE_SERVICE_ACCOUNT")
	serviceAccountFile := os.Getenv("GCE_SERVICE_ACCOUNT_FILE")

	if project == "" {
		return nil, fmt.Errorf("GCE_PROJECT environment variable not set")
	}

	// Service account can be provided as JSON string or file path
	var serviceAccountJSON string
	if serviceAccount != "" {
		serviceAccountJSON = serviceAccount
	} else if serviceAccountFile != "" {
		// Read service account file
		content, err := os.ReadFile(serviceAccountFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read GCE_SERVICE_ACCOUNT_FILE: %v", err)
		}
		serviceAccountJSON = string(content)
	}

	provider := &googleclouddns.Provider{
		Project:            project,
		ServiceAccountJSON: serviceAccountJSON,
	}

	return provider, nil
}
