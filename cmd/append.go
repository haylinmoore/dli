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
	appendTTL int
)

var appendCmd = &cobra.Command{
	Use:   "append <record_type> <name> <value>",
	Short: "Append a DNS record",
	Long:  "Append a DNS record to the specified zone using the specified provider. This adds the record without replacing existing ones.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if zone == "" {
			fmt.Printf("Error: --zone flag is required for append command\n")
			os.Exit(1)
		}

		recordType := args[0]
		name := args[1]
		value := args[2]

		fmt.Printf("Appending %s record: %s -> %s in zone %s using provider %s\n",
			recordType, name, value, zone, provider)

		// Create DNS record
		recordTTL := 300 // Default TTL of 5 minutes
		if appendTTL > 0 {
			recordTTL = appendTTL
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

		// Append the record
		ctx := context.Background()
		appendedRecords, err := providerInstance.AppendRecords(ctx, zone, []libdns.Record{record})
		if err != nil {
			fmt.Printf("Error appending record: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully appended record. Appended records: %+v\n", appendedRecords)
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
	appendCmd.Flags().IntVar(&appendTTL, "ttl", 0, "TTL in seconds (default: 300)")
}
