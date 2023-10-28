package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/tejiriaustin/slabmark-api/consumer"
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/events"
)

// apiCmd represents the api command
var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Starts slab-marks consumer service",
	Long:  ``,
	Run:   startListener,
}

func init() {
	rootCmd.AddCommand(listenerCmd)
}

func startListener(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	config := setListenerEnvironment()

	redis, err := database.NewRedisClient(config.GetAsString(env.RedisDsn), config.GetAsString(env.RedisPass), config.GetAsString(env.RedisPort))
	if err != nil {
		panic("Couldn't connect to redis dsn: " + err.Error())
	}
	defer func() {
		_ = redis.Disconnect(context.TODO())
	}()

	slabmarkListeners := consumer.NewConsumer().
		SetHandler(events.AccountAddedActivity, events.ForgotPasswordEventHandler())

	slabmarkListeners.ListenAndServe(ctx, redis)
}

func setListenerEnvironment() env.Environment {
	staticEnvironment := env.NewEnvironment()

	staticEnvironment.
		SetEnv(env.RedisDsn, env.MustGetEnv(env.RedisDsn)).
		SetEnv(env.RedisPort, env.MustGetEnv(env.RedisPort)).
		SetEnv(env.RedisPass, env.MustGetEnv(env.RedisPass)).
		SetEnv(env.MongoDsn, env.MustGetEnv(env.MongoDsn)).
		SetEnv(env.MongoDbName, env.MustGetEnv(env.MongoDbName)).
		SetEnv(env.MailjetApikeyPrivate, env.MustGetEnv(env.MailjetApikeyPrivate)).
		SetEnv(env.MailjetApikeyPublic, env.MustGetEnv(env.MailjetApikeyPublic))

	return staticEnvironment
}
