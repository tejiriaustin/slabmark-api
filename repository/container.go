package repository

import (
	"context"
	"fmt"
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Container struct {
		DB                 database.Database
		AccountsRepository *AccountsRepository
		LabRepository      *LabRepository
		StoreRepository    *StoreRepository
	}
	shared struct {
		collection database.Collection
		conf       env.Environment
	}
)

func NewContainer(db database.Database, conf env.Environment) *Container {
	return &Container{
		DB:                 db,
		AccountsRepository: NewAccountsRepository(db.GetCollection(fmt.Sprintf("%v.accounts", dbNameSpace)), conf),
		LabRepository:      NewLabRepository(db.GetCollection(fmt.Sprintf("%v.lab", dbNameSpace)), conf),
		StoreRepository:    NewStoreRepository(db.GetCollection(fmt.Sprintf("%v.lab", dbNameSpace)), conf),
	}
}

func (s shared) Count(ctx context.Context, queryFilter *QueryFilter) (int64, error) {
	return s.collection.CountDocuments(ctx, queryFilter.GetFilters())
}

func (s shared) DeleteMany(ctx context.Context, queryFilter *QueryFilter) error {
	_, err := s.collection.DeleteMany(ctx, queryFilter.GetFilters())
	return err
}

func (s shared) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s shared) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	//TODO implement me
	panic("implement me")
}

func (s shared) FindOne(ctx context.Context, filter interface{}, objects int, opts ...*options.FindOneOptions) *mongo.SingleResult {
	//TODO implement me
	panic("implement me")
}

func (s shared) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	//TODO implement me
	panic("implement me")
}

func (s shared) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s shared) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s shared) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}
