package events

import (
	"context"
	"fmt"

	"github.com/tejiriaustin/slabmark-api/consumer"
)

const (
	AccountAddedActivity = "ACTIVITY.ACCOUNT_ADDED"
)

func ForgotPasswordEventHandler() consumer.Handler {
	return func(ctx context.Context, msg consumer.Message) error {
		fmt.Println("Forgot Password")
		return nil
	}
}
