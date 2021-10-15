package restaurantservices

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
)

type UpdateRestaurantInterface interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
	Delete(ctx context.Context, id int) error
}

type UpdateRestaurantService struct {
	store     UpdateRestaurantInterface
	findStore *FindRestaurantService
}

func NewUpdateRestaurantService(conn *restaurantrepo.SqlConn) *UpdateRestaurantService {
	findStore := NewFindRestaurantService(conn)
	return &UpdateRestaurantService{
		store:     conn,
		findStore: findStore,
	}
}

func (r *UpdateRestaurantService) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	result, err := r.findStore.FindByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	if result.Status == 0 {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), nil)
	}
	err = r.store.Update(ctx, id, data)
	if err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	return err
}

func (r *UpdateRestaurantService) Delete(ctx context.Context, id int) error {
	_, err := r.findStore.FindByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	err = r.store.Delete(ctx, id)
	if err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	return err
}
