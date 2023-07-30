package repository

import (
	"context"
	"fmt"
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
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
		conf       *env.Environment
	}
)

func NewContainer(db database.Database, conf *env.Environment) *Container {
	return &Container{
		DB:                 db,
		AccountsRepository: (*AccountsRepository)(NewRepository(db.GetCollection(fmt.Sprintf("%v.accounts", dbNameSpace)), conf)),
		LabRepository:      (*LabRepository)(NewRepository(db.GetCollection(fmt.Sprintf("%v.lab", dbNameSpace)), conf)),
		StoreRepository:    (*StoreRepository)(NewRepository(db.GetCollection(fmt.Sprintf("%v.store", dbNameSpace)), conf)),
	}
}

func (s shared) Count(ctx context.Context, queryFilter *QueryFilter) (int64, error) {
	return s.collection.CountDocuments(ctx, queryFilter.GetFilters())
}

func (s shared) DeleteMany(ctx context.Context, queryFilter *QueryFilter) error {
	_, err := s.collection.DeleteMany(ctx, queryFilter.GetFilters())
	return err
}
