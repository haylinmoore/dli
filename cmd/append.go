package cmd

import (
	"github.com/spf13/cobra"
)

var appendCmd = &cobra.Command{
	Use:   "append [record_type] [name] [data]",
	Short: "Append DNS records",
	Long:  "Append DNS records to the specified zone using the specified provider. Use subcommands for specific record types, or provide record_type, name, and data directly for generic records.",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		// If no args provided, show help
		if len(args) == 0 {
			cmd.Help()
			return
		}

		// If exactly 3 args provided, treat as generic record
		if len(args) == 3 {
			recordType := args[0]
			name := args[1]
			data := args[2]

			executeRecordAppend(recordType, name, data)
			return
		}

		// Otherwise, show help for incorrect usage
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
	setupRecordCommand(appendCmd)
}
