package cmd

import (
	"context"
	"fmt"

	"dli/providers"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

var listZonesCmd = &cobra.Command{
	Use:   "list-zones",
	Short: "List DNS zones",
	Long:  "List all DNS zones available for the specified provider.",
	Run: func(cmd *cobra.Command, args []string) {
		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			OutputError("Failed to get provider", err)
			return
		}

		zoneLister, ok := providerInstance.(libdns.ZoneLister)
		if !ok {
			OutputError("Zone listing not supported", fmt.Errorf("provider %s does not support zone listing", provider))
			return
		}

		ctx := context.Background()
		zones, err := zoneLister.ListZones(ctx)
		if err != nil {
			OutputError("Failed to list zones", err)
			return
		}

		var zoneList []string
		for _, zone := range zones {
			zoneList = append(zoneList, zone.Name)
		}

		result := map[string]interface{}{
			"provider":   provider,
			"zone_count": len(zones),
			"zones":      zoneList,
		}
		OutputSuccess("Successfully retrieved zones", result)
	},
}

func init() {
	rootCmd.AddCommand(listZonesCmd)
}
