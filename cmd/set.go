package cmd

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [record_type] [name] [data]",
	Short: "Set DNS records",
	Long:  "Set DNS records in the specified zone using the specified provider. Use subcommands for specific record types, or provide record_type, name, and data directly for generic records.",
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

			executeRecordSet(recordType, name, data)
			return
		}

		// Otherwise, show help for incorrect usage
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setupRecordCommand(setCmd)
}
