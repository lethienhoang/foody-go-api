package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *sqlConn) Create (ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
