package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"dli/providers"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

var (
	ttl int
)

var setCmd = &cobra.Command{
	Use:   "set <record_type> <name> <value>",
	Short: "Set a DNS record",
	Long:  "Set a DNS record in the specified zone using the specified provider.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if zone == "" {
			fmt.Printf("Error: --zone flag is required for set command\n")
			os.Exit(1)
		}

		recordType := args[0]
		name := args[1]
		value := args[2]

		fmt.Printf("Setting %s record: %s -> %s in zone %s using provider %s\n",
			recordType, name, value, zone, provider)

		// Create DNS record
		recordTTL := 300 // Default TTL of 5 minutes
		if ttl > 0 {
			recordTTL = ttl
		}

		record := libdns.RR{
			Type: recordType,
			Name: name,
			Data: value,
			TTL:  time.Duration(recordTTL) * time.Second,
		}

		// Get provider instance
		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			fmt.Printf("Error getting provider: %v\n", err)
			os.Exit(1)
		}

		// Set the record
		ctx := context.Background()
		updatedRecords, err := providerInstance.SetRecords(ctx, zone, []libdns.Record{record})
		if err != nil {
			fmt.Printf("Error setting record: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully set record. Updated records: %+v\n", updatedRecords)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().IntVar(&ttl, "ttl", 0, "TTL in seconds (default: 300)")
}
