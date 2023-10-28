package publisher

import (
	"context"
	"sync"

	"github.com/tejiriaustin/slabmark-api/database"
)

type (
	Publisher struct {
		m      sync.Mutex
		client *database.RedisClient
	}

	PublishInterface interface {
		Publish(ctx context.Context, key string, message map[string]interface{}) error
	}
)

func newPublisher(client *database.RedisClient) *Publisher {
	return &Publisher{
		m:      sync.Mutex{},
		client: client,
	}
}
func NewPublisher(client *database.RedisClient) PublishInterface {
	return newPublisher(client)
}

func (p *Publisher) Publish(ctx context.Context, key string, message map[string]interface{}) error {
	return p.client.Publish(ctx, key, message)
}
