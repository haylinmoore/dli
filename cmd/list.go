package cmd

import (
	"context"
	"fmt"
	"os"

	"dli/providers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS records",
	Long:  "List all DNS records in the specified zone using the specified provider.",
	Run: func(cmd *cobra.Command, args []string) {
		if zone == "" {
			fmt.Printf("Error: --zone flag is required for list command\n")
			fmt.Printf("Use 'dli list-zones' to see available zones\n")
			os.Exit(1)
		}

		fmt.Printf("Listing DNS records in zone %s using provider %s\n", zone, provider)

		// Get provider instance
		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			fmt.Printf("Error getting provider: %v\n", err)
			os.Exit(1)
		}

		// Get all records
		ctx := context.Background()
		records, err := providerInstance.GetRecords(ctx, zone)
		if err != nil {
			fmt.Printf("Error getting records: %v\n", err)
			os.Exit(1)
		}

		if len(records) == 0 {
			fmt.Printf("No records found in zone %s\n", zone)
			return
		}

		fmt.Printf("Found %d records:\n", len(records))
		for _, record := range records {
			rr := record.RR()
			fmt.Printf("%-20s %-6s %-10s %s\n", rr.Name, rr.Type, rr.TTL, rr.Data)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
