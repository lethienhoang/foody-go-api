package restaurantservices

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
)

type FindRestaurantInterface interface {
	FindByCondition(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type FindRestaurantService struct {
	store FindRestaurantInterface
}

func NewFindRestaurantService(conn *restaurantrepo.SqlConn) *FindRestaurantService {
	return &FindRestaurantService{
		store: conn,
	}
}

func (r *FindRestaurantService) FindByCondition(ctx context.Context, condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	result, err := r.store.FindByCondition(ctx, condition, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	return result, nil
}
