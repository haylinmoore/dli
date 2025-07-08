package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	provider string
	zone     string
)

var rootCmd = &cobra.Command{
	Use:   "dli",
	Short: "DNS CLI tool for managing DNS records across multiple providers",
	Long:  "A CLI tool that allows you to manage DNS records across various DNS providers using environment variables for authentication.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&provider, "provider", "", "DNS provider (e.g., bunny, cloudflare, route53)")
	rootCmd.PersistentFlags().StringVar(&zone, "zone", "", "DNS zone to manage")
	rootCmd.MarkPersistentFlagRequired("provider")
	// Zone is not required for list-zones command
}
