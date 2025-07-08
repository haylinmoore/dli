package providers

import (
	"fmt"
	"os"

	"github.com/libdns/azure"
)

func getAzureProvider() (Provider, error) {
	subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroup := os.Getenv("AZURE_RESOURCE_GROUP")
	tenantId := os.Getenv("AZURE_TENANT_ID")
	clientId := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")

	if subscriptionId == "" {
		return nil, fmt.Errorf("AZURE_SUBSCRIPTION_ID environment variable not set")
	}
	if resourceGroup == "" {
		return nil, fmt.Errorf("AZURE_RESOURCE_GROUP environment variable not set")
	}

	provider := &azure.Provider{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroup,
		TenantId:          tenantId,
		ClientId:          clientId,
		ClientSecret:      clientSecret,
	}

	return provider, nil
}
