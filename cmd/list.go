package cmd

import (
	"context"
	"fmt"

	"dli/providers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS records",
	Long:  "List all DNS records in the specified zone using the specified provider.",
	Run: func(cmd *cobra.Command, args []string) {
		if zone == "" {
			OutputError("Zone required", fmt.Errorf("--zone flag is required for list command"))
			return
		}

		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			OutputError("Failed to get provider", err)
			return
		}

		ctx := context.Background()
		records, err := providerInstance.GetRecords(ctx, zone)
		if err != nil {
			OutputError("Failed to get records", err)
			return
		}

		var recordList []map[string]interface{}
		for _, record := range records {
			rr := record.RR()
			recordList = append(recordList, map[string]interface{}{
				"name": rr.Name,
				"type": rr.Type,
				"ttl":  int(rr.TTL.Seconds()),
				"data": rr.Data,
			})
		}

		result := map[string]interface{}{
			"zone":         zone,
			"provider":     provider,
			"record_count": len(records),
			"records":      recordList,
		}
		OutputSuccess("Successfully retrieved records", result)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
