package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Client struct {
	c  *mongo.Client
	DB *mongo.Database
}

type (
	Database interface {
		Disconnect(ctx context.Context) error
		GetCollection(name string, opts ...*options.CollectionOptions) Collection
	}
	Collection interface {
		CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
		DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
		FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
		FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
		InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	}
)

func Connect(dsn, dbName string, opts ...*options.ClientOptions) (*Client, error) {
	opts = append(opts, options.Client().ApplyURI(dsn))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mClient, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}

	err = mClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	database := mClient.Database(dbName)
	return &Client{
		DB: database,
		c:  mClient,
	}, nil
}

var _ Database = &Client{}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.c.Disconnect(ctx)
}

func (c *Client) GetCollection(name string, opts ...*options.CollectionOptions) Collection {
	return c.DB.Collection(name, opts...)
}
