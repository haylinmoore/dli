package cmd

import (
	"context"
	"fmt"
	"os"

	"dli/providers"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <record_type> <name> [value]",
	Short: "Delete a DNS record",
	Long:  "Delete a DNS record from the specified zone using the specified provider. If value is omitted, all records of the specified type and name will be deleted.",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		if zone == "" {
			fmt.Printf("Error: --zone flag is required for delete command\n")
			os.Exit(1)
		}

		recordType := args[0]
		name := args[1]
		var value string
		if len(args) == 3 {
			value = args[2]
		}

		// Get provider instance
		providerInstance, err := providers.GetProvider(provider)
		if err != nil {
			fmt.Printf("Error getting provider: %v\n", err)
			os.Exit(1)
		}

		ctx := context.Background()

		if value != "" {
			// Delete specific record
			fmt.Printf("Deleting %s record: %s -> %s in zone %s using provider %s\n",
				recordType, name, value, zone, provider)

			record := libdns.RR{
				Type: recordType,
				Name: name,
				Data: value,
			}

			deletedRecords, err := providerInstance.DeleteRecords(ctx, zone, []libdns.Record{record})
			if err != nil {
				fmt.Printf("Error deleting record: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Successfully deleted specific record. Deleted records: %d\n", len(deletedRecords))
			for _, deletedRecord := range deletedRecords {
				rr := deletedRecord.RR()
				fmt.Printf("Deleted: %s %s %s\n", rr.Name, rr.Type, rr.Data)
			}
		} else {
			// Delete all records of this type and name
			fmt.Printf("Deleting all %s records for %s in zone %s using provider %s\n",
				recordType, name, zone, provider)

			// First, get all records to find matching ones
			allRecords, err := providerInstance.GetRecords(ctx, zone)
			if err != nil {
				fmt.Printf("Error getting records: %v\n", err)
				os.Exit(1)
			}

			// Filter records matching type and name
			var recordsToDelete []libdns.Record
			for _, record := range allRecords {
				rr := record.RR()
				if rr.Type == recordType && rr.Name == name {
					recordsToDelete = append(recordsToDelete, record)
				}
			}

			if len(recordsToDelete) == 0 {
				fmt.Printf("No matching records found for %s %s\n", recordType, name)
				return
			}

			deletedRecords, err := providerInstance.DeleteRecords(ctx, zone, recordsToDelete)
			if err != nil {
				fmt.Printf("Error deleting records: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Successfully deleted all matching records. Deleted records: %d\n", len(deletedRecords))
			for _, deletedRecord := range deletedRecords {
				rr := deletedRecord.RR()
				fmt.Printf("Deleted: %s %s %s\n", rr.Name, rr.Type, rr.Data)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
