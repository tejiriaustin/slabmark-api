package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/server"
	"github.com/tejiriaustin/slabmark-api/services"
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
	ctx := context.Background()

	config := setApiEnvironment()

	dbConn, err := database.NewMongoDbClient().Connect(config.GetAsString(env.MONGO_DSN), config.GetAsString(env.MONGO_DB_NAME))
	if err != nil {
		panic("Couldn't connect to mongo dsn: " + err.Error())
	}
	defer func() {
		_ = dbConn.Disconnect(context.TODO())
	}()

	redis, err := database.NewRedisClient(config.GetAsString(env.REDIS_DSN), config.GetAsString(env.REDIS_PORT))
	if err != nil {
		panic("Couldn't connect to redis dsn: " + err.Error())
	}
	defer func() {
		_ = redis.Disconnect(context.TODO())
	}()

	rc := repository.NewContainer(nil, config)

	sc := services.NewService(config)

	server.Start(ctx, sc, rc)

}

func setApiEnvironment() *env.Environment {
	staticEnvironment := env.NewEnvironment()

	staticEnvironment.
		SetEnv(env.EnvPort, env.GetEnv(env.EnvPort, "8080")).
		SetEnv(env.MONGO_DSN, env.MustGetEnv(env.MONGO_DSN)).
		SetEnv(env.REDIS_DSN, env.MustGetEnv(env.REDIS_DSN)).
		SetEnv(env.MONGO_DB_NAME, env.MustGetEnv(env.MONGO_DB_NAME))

	return staticEnvironment
}
