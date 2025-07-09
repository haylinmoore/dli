package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	provider string
	zone     string
	envFile  string
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

// loadEnvFromFile loads environment variables from a file
// Variables in the file override system environment variables
func loadEnvFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// Don't exit, just warn - the user might want to proceed without the env file
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split on first = to handle values with = in them
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		// Set the environment variable (overrides system env)
		os.Setenv(key, value)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&provider, "provider", "", "DNS provider (e.g., bunny, cloudflare, route53)")
	rootCmd.PersistentFlags().StringVar(&zone, "zone", "", "DNS zone to manage")
	rootCmd.PersistentFlags().StringVar(&envFile, "env", "", "Path to environment file (variables in file override system environment)")
	rootCmd.MarkPersistentFlagRequired("provider")
	// Zone is not required for list-zones command

	// Load environment variables from file if specified
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if envFile != "" {
			loadEnvFromFile(envFile)
		}
	}
}
