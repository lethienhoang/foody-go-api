package restaurantservices

import (
	"context"
	"errors"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
)

type CreateRestaurantInterface interface {
	Create (ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type CreateRestaurantService struct {
	store CreateRestaurantInterface
}

func NewCreateRestaurantService(conn *restaurantrepo.SqlConn) *CreateRestaurantService {
	return &CreateRestaurantService{
		store: conn,
	}
}

func(r *CreateRestaurantService) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return common.ErrInvalidRequest(errors.New("restaurant name can not be blank"))
	}

	err := r.store.Create(ctx, data)

	return err
}
