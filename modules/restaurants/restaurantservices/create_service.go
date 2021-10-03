package restaurantservices

import (
	"context"
	"errors"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

type CreateRestaurantInterface interface {
	Create (ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type CreateRestaurantService struct {
	store CreateRestaurantInterface
}

func NewCreateRestaurantService(store CreateRestaurantInterface) *CreateRestaurantService {
	return &CreateRestaurantService{
		store: store,
	}
}

func(r *CreateRestaurantService) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("restaurant name can not be blank")
	}

	err := r.store.Create(ctx, data)

	return err
}
