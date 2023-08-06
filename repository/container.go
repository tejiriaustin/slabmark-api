package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/tejiriaustin/slabmark-api/database"
	"github.com/tejiriaustin/slabmark-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	dbnamespace = "slabmark_api_db_cluster"
)

type Container struct {
	AccountsRepo *Repository[models.Account]
	LabsRepo     *Repository[models.LabReading]
	StoreRepo    *Repository[models.StoreItem]
}

func NewRepositoryContainer(dbConn *database.Client) *Container {
	log.Println(" building repository container...")

	return &Container{
		AccountsRepo: NewRepository[models.Account](dbConn.GetCollection(fmt.Sprintf("%v.accounts", dbnamespace))),
		LabsRepo:     NewRepository[models.LabReading](dbConn.GetCollection(fmt.Sprintf("%v.labsrepo", dbnamespace))),
		StoreRepo:    NewRepository[models.StoreItem](dbConn.GetCollection(fmt.Sprintf("%v.store", dbnamespace))),
	}
}

type Repository[T models.SharedInterface] struct {
	dbCollection database.Collection
}

func NewRepository[T models.SharedInterface](dbCollection database.Collection) *Repository[T] {
	return &Repository[T]{dbCollection: dbCollection}
}

func (r *Repository[T]) Create(ctx context.Context, data T) (T, error) {
	data.Initialize(primitive.NewObjectID(), time.Now())

	res, err := r.dbCollection.InsertOne(ctx, data)
	if err != nil {
		return data, errors.New("failed to insert one")
	}
	data.SetID(res.InsertedID.(primitive.ObjectID))
	return data, nil
}

func (r *Repository[T]) FindOne(ctx context.Context, queryFilter *QueryFilter, projection *QueryProjection, findOneOptions ...*options.FindOneOptions) (T, error) {
	var data T

	opts := &options.FindOneOptions{}
	if projection == nil {
		opts.Projection = projection.GetProjection()
		data.SetUsedProjection(true)
	}

	err := r.dbCollection.FindOne(ctx, queryFilter.GetFilters(), findOneOptions...).Decode(data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return data, errors.New("no documents found")
		}
		return data, errors.New("failed find One")
	}
	return data, nil
}

func (r *Repository[T]) Update(ctx context.Context, dataObject T) (T, error) {

	//if dataObject == nil {
	//	return dataObject, errors.New("dataObject can't be nil")
	//}
	if dataObject.DidUseProjection() {
		return dataObject, errors.New("can't Update Document That Was Queried With A Projection - Some Fields May Be Lost")
	}

	dataObject.SetUpdatedAt()
	queryFilter := NewQueryFilter().AddFilter("_id", dataObject.GetId())
	res := r.dbCollection.FindOneAndReplace(ctx, queryFilter.GetFilters(), dataObject)

	if res.Err() != nil {
		return dataObject, errors.New(fmt.Sprintf("Updated Failed with error: %s", res.Err()))
	}

	return dataObject, nil
}

// findPaginated searches for document that matches the provided filters.
// paginatorOptions control CurrentPage and PerPage value.
// If projection is nil, all fields are returned.
// sort should be a bson.D - eg: bson.D{bson.E{Key: "_id", Value: -1}, bson.E{Key: "another, Value: "value"}}
// findPaginated will return the Mongo Cursor in the paginatedResult struct.
func (r *Repository[T]) findPaginated(ctx context.Context, pageOptions paginatorOptions, filters *QueryFilter, projection *QueryProjection, sort *QuerySort) (*paginatedResult, error) {
	if sort == nil {
		sort = NewDefaultQuerySort()
	}

	paginator := newPaginator(pageOptions)
	paginator.setOffset()
	opts := &options.FindOptions{
		Skip:  &paginator.Offset,
		Limit: &paginator.PerPage,
		Sort:  sort.GetSort(),
	}

	if projection != nil {
		opts.Projection = projection.GetProjection()
	}

	totalRows, err := r.dbCollection.CountDocuments(ctx, filters.GetFilters())
	if err != nil {
		return nil, err
	}
	paginator.TotalRows = totalRows

	cur, err := r.dbCollection.Find(ctx, filters.GetFilters(), opts)
	if err != nil {
		return nil, err
	}

	paginator.setTotalPages()
	paginator.setPrevPage()
	paginator.setNextPage()
	return &paginatedResult{Cursor: cur, Paginator: paginator}, nil
}

func (r *Repository[T]) Paginate(
	ctx context.Context,
	filters *QueryFilter,
	page, perPage int64,
	projection *QueryProjection,
	sort *QuerySort) ([]T, *Paginator, error) {

	var dataObjects []T
	po := paginatorOptions{
		Page:    page,
		PerPage: perPage,
	}
	res, err := r.findPaginated(ctx, po, filters, projection, sort)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return dataObjects, nil, errors.New("no data Objects Found")
		}
		return dataObjects, nil, errors.New("pagination Failed")
	}

	defer func(Cursor *mongo.Cursor, ctx context.Context) {
		err := Cursor.Close(ctx)
		if err != nil {
			log.Println("Cursor.Close failed to close cursor")
		}
	}(res.Cursor, ctx)

	for res.Cursor.Next(ctx) {
		var dataObject T
		err := res.Cursor.Decode(&dataObject)
		if err == nil {
			dataObjects = append(dataObjects, dataObject)
			continue
		}
		return dataObjects, nil, errors.New("failed to decode")
	}
	return dataObjects, res.Paginator, nil
}
