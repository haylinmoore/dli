package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [record_type] [name] [data]",
	Short: "Delete DNS records",
	Long:  "Delete DNS records from the specified zone using the specified provider. Use subcommands for specific record types, or provide record_type, name, and optionally data directly for generic records. If data is omitted, all records of the specified type and name will be deleted.",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		// If no args provided, show help
		if len(args) == 0 {
			cmd.Help()
			return
		}

		// If 2 or 3 args provided, treat as generic record delete
		if len(args) == 2 || len(args) == 3 {
			recordType := args[0]
			name := args[1]
			var data string
			if len(args) == 3 {
				data = args[2]
			}

			executeRecordDelete(recordType, name, data)
			return
		}

		// Otherwise, show help for incorrect usage
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
