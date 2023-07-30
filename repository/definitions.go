package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
)

const (
	dbNameSpace = "slabmark-api-collection"
)

type T struct {
}

type (
	Repository[T any] struct {
		shared
	}
	AccountsRepository struct {
		shared
	}
	LabRepository struct {
		shared
	}
	StoreRepository struct {
		shared
	}
)

type IRepository interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

func NewRepository[T any](dbCollection database.Collection, conf *env.Environment) *Repository[T] {
	return &Repository[T]{
		shared: shared{
			collection: dbCollection,
			conf:       conf,
		},
	}
}

func (s Repository[T]) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (s Repository[T]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	//TODO implement me
	panic("implement me")
}

func (s Repository[T]) FindOne(ctx context.Context, filter interface{}, objects int, opts ...*options.FindOneOptions) *mongo.SingleResult {
	//TODO implement me
	panic("implement me")
}

func (s Repository[T]) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	//TODO implement me
	panic("implement me")
}

func (s Repository[T]) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s Repository[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s shared) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}
