package restaurantservices

import (
	"context"
	"errors"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

type UpdateRestaurantInterface interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
	FindByCondition(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type UpdateRestaurantService struct {
	store UpdateRestaurantInterface
}

func NewUpdateRestaurantService(store UpdateRestaurantInterface) *UpdateRestaurantService {
	return &UpdateRestaurantService{
		store: store,
	}
}

func (r *UpdateRestaurantService) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	result, err := r.store.FindByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if result.Status == 0 {
		errors.New("data is deleted")
	}
	err = r.store.Update(ctx, id, data)
	if err != nil {
		return err
	}

	return err
}
