package restaurantservices

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

type ListRestaurantInterface interface {
	ListByCondition(ctx context.Context, conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantService struct {
	store ListRestaurantInterface
}

func NewListRestaurantService(store ListRestaurantInterface) *ListRestaurantService {
	return &ListRestaurantService{store: store}
}

func (l *ListRestaurantService) ListRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := l.store.ListByCondition(ctx, conditions, filter, paging)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	return result, nil
}
