package cmd

import "github.com/spf13/cobra"

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts slab-marks api",
	Long:  ``,
	Run:   startApi,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startApi(cmd *cobra.Command, args []string) {

	// Build the server

	// Build the repository

	// Build the services

	// Build notifications

	// Start Database Connection

	// Run the Server
}
