package consumer

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/tejiriaustin/slabmark-api/database"
)

const (
	defaultMaxWorkers uint = 10
)

type (
	consumer struct {
		maxWorkerChans uint
		refreshTime    int
		handlers       map[string]Handler
	}

	Message struct {
		Key  string `json:"key"`
		Body string `json:"body"`
	}

	Handler func(ctx context.Context, msg Message) error
	Options func(*consumer)
)

func newConsumer() *consumer {
	return &consumer{
		maxWorkerChans: defaultMaxWorkers,
		refreshTime:    0,
		handlers:       make(map[string]Handler),
	}
}

func NewConsumer(opts ...Options) *consumer {
	l := newConsumer()

	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *consumer) SetHandler(key string, handler Handler) *consumer {
	l.handlers[strings.ToUpper(key)] = handler
	return l
}

func (l *consumer) ListenAndServe(ctx context.Context, redisClient *database.RedisClient) {
	log.Print("initializing  consumer...")

	workerChannels := make(chan Message, l.maxWorkerChans)
	defer func() {
		close(workerChannels)
	}()

	pubsub := redisClient.Subscribe(ctx)
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			zap.L().Error("failed to close redis subscription connection", zap.Error(err))
		}
	}(pubsub)

	zap.L().Info("starting workers", zap.Uint("maxWorkers", l.maxWorkerChans))

	for i := uint(0); i < l.maxWorkerChans; i++ {
		go l.worker(ctx, workerChannels)
	}

	for {
		zap.L().Info("pulling messages...")
		message, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			zap.L().Error("failed to receive message", zap.Error(err))
		}
		msg := new(Message)
		err = json.Unmarshal([]byte(message.Payload), msg)
		if err != nil {
			return
		}

		workerChannels <- *msg
	}
}

// worker listens on the msgChan for incoming messages.
// as messages become available over the channel, it looks over the map of configured handlers and routes messages by the key.
// If no handlers are configured and a default handler has been set, the message is sent there.
// else it logs and continues.
func (l *consumer) worker(ctx context.Context, msgChan <-chan Message) {
	for msg := range msgChan {
		_ = l.dispatcher(ctx, msg)
	}
}

func (l *consumer) dispatcher(ctx context.Context, message Message) error {
	msg := new(Message)

	err := json.Unmarshal([]byte(message.Body), msg)
	if err != nil {
		return err
	}

	handlerFunc := l.handlers[message.Key]

	if handlerFunc == nil {
		zap.L().Info("handlerfunc is nil", zap.Error(err), zap.String("message", message.Key))
		return nil
	}

	if err := handlerFunc(ctx, message); err != nil {
		zap.L().Error("failed to handle message", zap.Error(err), zap.String("message", message.Key))
		return err
	}

	return nil
}
