package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *sqlConn) Update (ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id = ?",id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

