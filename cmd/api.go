package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tejiriaustin/slabmark-api/env"
)

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

func setApiEnvironment() *env.Environment {
	staticEnvironment := env.NewEnvironment()

	staticEnvironment.
		SetEnv(env.EnvPort, env.GetEnv(env.EnvPort, "8080")).
		SetEnv(env.MONGO_DSN, env.MustGetEnv(env.MONGO_DSN)).
		SetEnv(env.REDIS_DSN, env.MustGetEnv(env.REDIS_DSN))

	return staticEnvironment
}
