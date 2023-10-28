package services

import (
	"context"

	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type (
	ActivityService struct {
		conf *env.Environment
	}

	ActivityListFilters struct {
		Query string // for partial free hand lookups
	}
	ListActivityInput struct {
		Pager
		Projection *repository.QueryProjection
		Sort       *repository.QuerySort
		Filters    ActivityListFilters
	}
)

func NewActivityService(conf *env.Environment) *ActivityService {
	return &ActivityService{conf: conf}
}

func (s *ActivityService) ListActivities(
	ctx context.Context,
	input ListActivityInput,
	activityRepo *repository.Repository[models.Activity],
) ([]models.Activity, *repository.Paginator, error) {

	filter := repository.NewQueryFilter()

	if input.Filters.Query != "" {
		freeHandFilters := []map[string]interface{}{
			{"first_name": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"last_name": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"full_name": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"phone": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"email": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"username": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
		}
		filter.AddFilter("$or", freeHandFilters)
	}

	account, paginator, err := activityRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return account, paginator, nil

}
