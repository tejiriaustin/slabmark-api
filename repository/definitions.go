package repository

import (
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/env"
)

const (
	dbNameSpace = "slabmark-api-collection"
)

type (
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

func NewAccountsRepository(dbCollection database.Collection, conf env.Environment) *AccountsRepository {
	return &AccountsRepository{
		shared: shared{
			collection: dbCollection,
			conf:       conf,
		},
	}
}

func NewLabRepository(dbCollection database.Collection, conf env.Environment) *LabRepository {
	return &LabRepository{
		shared: shared{
			collection: dbCollection,
			conf:       conf,
		},
	}
}

func NewStoreRepository(dbCollection database.Collection, conf env.Environment) *StoreRepository {
	return &StoreRepository{
		shared: shared{
			collection: dbCollection,
			conf:       conf,
		},
	}
}
