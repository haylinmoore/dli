package cmd

import (
	"context"
	"fmt"
	"os"

	"dli/providers"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

var listZonesCmd = &cobra.Command{
	Use:   "list-zones",
	Short: "List DNS zones",
	Long:  "List all DNS zones available for the specified provider.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing DNS zones for provider %s\n", provider)

		// Get provider instance
		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			fmt.Printf("Error getting provider: %v\n", err)
			os.Exit(1)
		}

		// Check if provider supports zone listing
		zoneLister, ok := providerInstance.(libdns.ZoneLister)
		if !ok {
			fmt.Printf("Provider %s does not support zone listing\n", provider)
			os.Exit(1)
		}

		// List zones
		ctx := context.Background()
		zones, err := zoneLister.ListZones(ctx)
		if err != nil {
			fmt.Printf("Error listing zones: %v\n", err)
			os.Exit(1)
		}

		if len(zones) == 0 {
			fmt.Printf("No zones found for provider %s\n", provider)
			return
		}

		fmt.Printf("Found %d zones:\n", len(zones))
		for _, zone := range zones {
			fmt.Printf("- %s\n", zone.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listZonesCmd)
}
